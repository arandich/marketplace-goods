package grpc

import (
	"context"
	"github.com/arandich/marketplace-goods/internal/config"
	"github.com/arandich/marketplace-goods/internal/model"
	"github.com/arandich/marketplace-goods/pkg/metrics"
	pb "github.com/arandich/marketplace-proto/api/proto/services"
	loggerInterceptor "github.com/arandich/marketplace-sdk/interceptors/logger"
	recoveryInterceptor "github.com/arandich/marketplace-sdk/interceptors/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"net"
)

type Server struct {
	logger *zerolog.Logger
	cfg    config.GrpcConfig
}

func New(ctx context.Context, cfg config.GrpcConfig) *Server {
	return &Server{
		cfg:    cfg,
		logger: zerolog.Ctx(ctx),
	}
}

func (s *Server) Start(ctx context.Context, lis net.Listener, services model.Services, promMetrics metrics.Metrics) (*grpc.Server, chan error) {
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(s.cfg.ConnectionTimeout),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     s.cfg.MaxConnectionIdle,
			MaxConnectionAge:      s.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: s.cfg.MaxConnectionAgeGrace,
			Timeout:               s.cfg.KeepAliveTimeout,
		}),

		grpc.ChainUnaryInterceptor(
			recoveryInterceptor.UnaryServerInterceptor(),
			loggerInterceptor.NewUnaryLoggerInterceptor(ctx),
			promMetrics.BaseMetrics.UnaryServerInterceptor(ctx),
		),
	}

	grpcSrv := grpc.NewServer(opts...)

	pb.RegisterGoodsServiceServer(grpcSrv, services.GoodsService)

	// Disable if not dev development
	reflection.Register(grpcSrv)

	errChan := make(chan error, 1)
	go func() {
		errChan <- grpcSrv.Serve(lis)
	}()

	return grpcSrv, errChan
}

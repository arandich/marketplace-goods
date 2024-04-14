package repository

import (
	"context"
	"github.com/arandich/marketplace-goods/internal/config"
	"github.com/arandich/marketplace-goods/internal/model"
	"github.com/arandich/marketplace-goods/pkg/metrics"
	pb "github.com/arandich/marketplace-proto/api/proto/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type GoodsRepository struct {
	pgPool      *pgxpool.Pool
	promMetrics metrics.Metrics
	logger      *zerolog.Logger
	cfg         config.Config
	clients     model.Clients
}

func NewGoodsRepository(ctx context.Context, pgPool *pgxpool.Pool, promMetrics metrics.Metrics, cfg config.Config, clients model.Clients) *GoodsRepository {
	return &GoodsRepository{
		pgPool:      pgPool,
		promMetrics: promMetrics,
		logger:      zerolog.Ctx(ctx),
		cfg:         cfg,
		clients:     clients,
	}
}

func (g GoodsRepository) GetGood(ctx context.Context, req *pb.GetGoodRequest) (*pb.GetGoodResponse, error) {
	return nil, nil
}

func (g GoodsRepository) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsResponse, error) {
	return nil, nil
}

func (g GoodsRepository) AddGood(ctx context.Context, req *pb.AddGoodRequest) (*pb.AddGoodResponse, error) {
	return nil, nil
}

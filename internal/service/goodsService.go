package service

import (
	"context"
	pb "github.com/arandich/marketplace-proto/api/proto/services"
)

type GoodsRepository interface {
	GetGood(ctx context.Context, req *pb.GetGoodRequest) (*pb.GetGoodResponse, error)
	GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsResponse, error)
	AddGood(ctx context.Context, req *pb.AddGoodRequest) (*pb.AddGoodResponse, error)
}

var _ GoodsRepository = (*GoodsService)(nil)

type GoodsService struct {
	pb.UnimplementedGoodsServiceServer
	repository GoodsRepository
}

func NewIdService(repository GoodsRepository) GoodsService {
	return GoodsService{
		repository: repository,
	}
}

func (s GoodsService) GetGood(ctx context.Context, req *pb.GetGoodRequest) (*pb.GetGoodResponse, error) {
	return s.repository.GetGood(ctx, req)
}

func (s GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsResponse, error) {
	return s.repository.GetGoods(ctx, req)
}

func (s GoodsService) AddGood(ctx context.Context, req *pb.AddGoodRequest) (*pb.AddGoodResponse, error) {
	return s.repository.AddGood(ctx, req)
}

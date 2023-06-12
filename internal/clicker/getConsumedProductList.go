package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) GetConsumedProductList(ctx context.Context, _ *pb.GetConsumedProductListRequest) (*pb.GetConsumedProductListResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	consumedProducts, err := s.Storage.GetConsumedProductList(token)
	if err != nil {
		return nil, err
	}
	return &pb.GetConsumedProductListResponse{Products: consumedProducts}, nil
}

package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) GetAvailableProductList(_ context.Context, _ *pb.GetAvailableProductListRequest) (*pb.GetAvailableProductListResponse, error) {
	availableProducts := s.Storage.GetAvailableProductList()
	return &pb.GetAvailableProductListResponse{Products: availableProducts}, nil
}

package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) ConsumeProduct(ctx context.Context, request *pb.ConsumeProductRequest) (*pb.ConsumeProductResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.Storage.ConsumeProduct(token, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ConsumeProductResponse{}, nil
}

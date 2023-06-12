package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) GetInt(ctx context.Context, request *pb.GetIntRequest) (*pb.GetIntResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	value, err := s.Storage.GetInt(token, request.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetIntResponse{Value: value}, nil
}

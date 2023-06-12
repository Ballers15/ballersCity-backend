package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) GetString(ctx context.Context, request *pb.GetStringRequest) (*pb.GetStringResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	value, err := s.Storage.GetString(token, request.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetStringResponse{Value: value}, nil
}

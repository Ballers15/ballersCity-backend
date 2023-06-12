package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) GetFloat(ctx context.Context, request *pb.GetFloatRequest) (*pb.GetFloatResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	value, err := s.Storage.GetFloat(token, request.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetFloatResponse{Value: value}, nil
}

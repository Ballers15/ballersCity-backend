package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) SetString(ctx context.Context, request *pb.SetStringRequest) (*pb.SetStringResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.Storage.SetString(token, request.Key, request.Value)
	if err != nil {
		return nil, err
	}
	return &pb.SetStringResponse{}, nil
}

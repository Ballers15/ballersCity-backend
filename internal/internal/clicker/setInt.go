package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) SetInt(ctx context.Context, request *pb.SetIntRequest) (*pb.SetIntResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.Storage.SetInt(token, request.Key, request.Value)
	if err != nil {
		return nil, err
	}
	return &pb.SetIntResponse{}, nil
}

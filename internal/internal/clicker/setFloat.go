package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) SetFloat(ctx context.Context, request *pb.SetFloatRequest) (*pb.SetFloatResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	err = s.Storage.SetFloat(token, request.Key, request.Value)
	if err != nil {
		return nil, err
	}
	return &pb.SetFloatResponse{}, nil
}

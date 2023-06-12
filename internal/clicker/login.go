package clicker

import (
	"context"
	"github.com/nostressdev/tvx/pb"
)

func (s *TvxServer) Login(_ context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.Storage.Login(request.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}

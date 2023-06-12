package inMemoryStorage

import (
	"github.com/nostressdev/nerrors"
	"github.com/nostressdev/tvx/pb"
)

func (s *IMStorage) GetConsumedProductList(accessToken string) (products []*pb.Product, err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return nil, nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	return s.UserConsumedProducts[userEmail], nil
}

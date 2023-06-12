package inMemoryStorage

import (
	"github.com/nostressdev/nerrors"
	"github.com/nostressdev/tvx/pb"
)

func (s *IMStorage) ConsumeProduct(accessToken string, id string) (err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	for _, pr := range s.Products {
		if pr.Id == id {
			if pr.IsConsumable {
				if s.UserConsumedProducts[userEmail] == nil {
					s.UserConsumedProducts[userEmail] = []*pb.Product{}
				}
				s.UserConsumedProducts[userEmail] = append(s.UserConsumedProducts[userEmail], pr)
				return nil
			}
			return nerrors.NotFound.New("product is not consumable")
		}
	}
	return nerrors.NotFound.New("no product with such id")
}

package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"github.com/nostressdev/tvx/pb"
)

func (s *PostgresStorage) GetConsumedProductList(accessToken string) (products []*pb.Product, err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return nil, nerrors.BadRequest.New("invalid access token")
	}

	userConsumedProductModels := []*UserConsumedProductModel{}
	s.db.Where("email = ?", userAccessToken.Email).Find(&userConsumedProductModels)
	consumedProductList := make([]*pb.Product, len(userConsumedProductModels))
	for i, product := range userConsumedProductModels {
		consumedProductList[i] = &pb.Product{
			Id:           product.PbId,
			PaymentType:  product.PaymentType,
			Price:        product.Price,
			IsConsumable: product.IsConsumable,
		}
	}
	return consumedProductList, nil
}

package postgresStorage

import (
	"github.com/nostressdev/nerrors"
)

func (s *PostgresStorage) ConsumeProduct(accessToken string, id string) (err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return nerrors.BadRequest.New("invalid access token")
	}

	var productModel ProductModel
	s.db.Where("pb_id = ?", id).First(&productModel)
	if productModel.PbId == "" {
		return nerrors.NotFound.New("no product with such id")
	}
	if productModel.IsConsumable {
		var userConsumedProductModel UserConsumedProductModel
		s.db.Where("email = ? AND pb_id = ?", userAccessToken.Email, productModel.PbId).First(&userConsumedProductModel)
		userConsumedProductModel.Id = id
		userConsumedProductModel.PbId = id
		userConsumedProductModel.IsConsumable = productModel.IsConsumable
		userConsumedProductModel.Price = productModel.Price
		userConsumedProductModel.PaymentType = productModel.PaymentType
		if userConsumedProductModel.Email == "" {
			userConsumedProductModel.Email = userAccessToken.Email
			s.db.Create(&userConsumedProductModel)
			return nil
		}
		s.db.Save(&userConsumedProductModel)
		return nil
	}
	return nerrors.NotFound.New("product is not consumable")
}

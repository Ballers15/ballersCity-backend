package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) SetFloat(accessToken string, key string, value float64) (err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsFloatModel UserSettingsFloatModel
	log.Println("SetFloat-0", "userSettingsFloatModel", userSettingsFloatModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsFloatModel)
	log.Println("SetFloat-1", "userSettingsFloatModel", userSettingsFloatModel)
	if userSettingsFloatModel.Email == "" {
		userSettingsFloatModel.Email = userAccessToken.Email
		userSettingsFloatModel.Key = key
		userSettingsFloatModel.Value = value
		s.db.Create(&userSettingsFloatModel)
		return nil
	}
	userSettingsFloatModel.Value = value
	s.db.Save(&userSettingsFloatModel)
	log.Println("SetFloat-2", "userSettingsFloatModel", userSettingsFloatModel)
	return nil
}

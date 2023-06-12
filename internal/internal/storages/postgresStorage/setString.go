package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) SetString(accessToken string, key string, value string) (err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsStringModel UserSettingsStringModel
	log.Println("SetString-0", "userSettingsStringModel", userSettingsStringModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsStringModel)
	log.Println("SetString-1", "userSettingsStringModel", userSettingsStringModel)
	if userSettingsStringModel.Email == "" {
		userSettingsStringModel.Email = userAccessToken.Email
		userSettingsStringModel.Key = key
		userSettingsStringModel.Value = value
		s.db.Create(&userSettingsStringModel)
		return nil
	}
	userSettingsStringModel.Value = value
	s.db.Save(&userSettingsStringModel)
	log.Println("SetString-2", "userSettingsStringModel", userSettingsStringModel)
	return nil
}

package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) SetInt(accessToken string, key string, value int64) (err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsIntModel UserSettingsIntModel
	log.Println("SetInt-0", "userSettingsIntModel", userSettingsIntModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsIntModel)
	log.Println("SetInt-1", "userSettingsIntModel", userSettingsIntModel)
	if userSettingsIntModel.Email == "" {
		userSettingsIntModel.Email = userAccessToken.Email
		userSettingsIntModel.Key = key
		userSettingsIntModel.Value = value
		s.db.Create(&userSettingsIntModel)
		return nil
	}
	userSettingsIntModel.Value = value
	s.db.Save(&userSettingsIntModel)
	log.Println("SetInt-2", "userSettingsIntModel", userSettingsIntModel)
	return nil
}

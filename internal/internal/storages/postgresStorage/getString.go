package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) GetString(accessToken string, key string) (value string, err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Where("access_token = ?", accessToken).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return "", nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsStringModel UserSettingsStringModel
	log.Println("GetString-0", "userSettingsStringModel", userSettingsStringModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsStringModel)
	log.Println("GetString-1", "userSettingsStringModel", userSettingsStringModel)
	if userSettingsStringModel.Email == "" {
		return "", nerrors.NotFound.New("no such key in user's int map")
	}
	return userSettingsStringModel.Value, nil
}

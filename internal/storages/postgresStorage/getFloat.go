package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) GetFloat(accessToken string, key string) (value float64, err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return 0, nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsFloatModel UserSettingsFloatModel
	log.Println("GetFloat-0", "userSettingsFloatModel", userSettingsFloatModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsFloatModel)
	log.Println("GetFloat-1", "userSettingsFloatModel", userSettingsFloatModel)
	if userSettingsFloatModel.Email == "" {
		return 0, nerrors.NotFound.New("no such key in user's int map")
	}
	return userSettingsFloatModel.Value, nil
}

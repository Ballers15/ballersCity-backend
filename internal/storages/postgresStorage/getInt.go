package postgresStorage

import (
	"github.com/nostressdev/nerrors"
	"log"
)

func (s *PostgresStorage) GetInt(accessToken string, key string) (value int64, err error) {
	var userAccessToken UserAccessTokensModel
	s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessToken)
	if userAccessToken.Email == "" {
		return 0, nerrors.BadRequest.New("invalid access token")
	}

	var userSettingsIntModel UserSettingsIntModel
	log.Println("GetInt-0", "userSettingsIntModel", userSettingsIntModel)
	s.db.Where("email = ? AND key = ?", userAccessToken.Email, key).First(&userSettingsIntModel)
	log.Println("GetInt-1", "userSettingsIntModel", userSettingsIntModel)
	if userSettingsIntModel.Email == "" {
		return 0, nerrors.NotFound.New("no such key in user's int map")
	}
	return userSettingsIntModel.Value, nil
}

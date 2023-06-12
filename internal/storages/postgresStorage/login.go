package postgresStorage

import (
	"fmt"
	"github.com/nostressdev/tvx/internal/storages"
)

func (s *PostgresStorage) Login(email string) (string, error) {
	accessToken, err := storages.GenerateOneTimePassword(32)
	if err != nil {
		return "", err
	}
	var userAccessTokensModel UserAccessTokensModel
	for {
		s.db.Model(UserAccessTokensModel{AccessToken: accessToken}).First(&userAccessTokensModel, "access_token = ?", accessToken)
		if userAccessTokensModel.Email == "" {
			userAccessTokensModel.AccessToken = accessToken
			userAccessTokensModel.Email = email
			s.db.Create(&userAccessTokensModel)
			fmt.Printf("Verification of user '%s' was completed! New token: '%s'\n", email, accessToken)
			return accessToken, nil
		}
		accessToken, err = storages.GenerateOneTimePassword(32)
		if err != nil {
			return "", err
		}
		userAccessTokensModel.Email = ""
	}
}

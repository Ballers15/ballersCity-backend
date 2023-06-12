package inMemoryStorage

import (
	"fmt"
	"github.com/nostressdev/nerrors"
	"github.com/nostressdev/tvx/internal/storages"
)

func (s *IMStorage) VerifyLogin(email string, code string) (accessToken string, err error) {
	if verificationCode, ok := s.EmailVerificationCodes[email]; !ok {
		return "", nerrors.BadRequest.New("this email is not awaiting verification")
	} else {
		if code == verificationCode {
			delete(s.EmailVerificationCodes, email)
			accessToken, err := storages.GenerateOneTimePassword(32)
			if err != nil {
				return "", err
			}
			for {
				if _, ok := s.UserAccessTokens[accessToken]; !ok {
					s.UserAccessTokens[accessToken] = email
					if _, ok := s.UserSettings[email]; !ok {
						s.UserSettings[email] = &UserSettings{StringMap: make(map[string]string),
							IntMap:   make(map[string]int64),
							FloatMap: make(map[string]float64)}
					}
					fmt.Printf("Verification of user '%s' was completed! New token: '%s'\n", email, accessToken)
					return accessToken, nil
				}
				accessToken, err = storages.GenerateOneTimePassword(32)
				if err != nil {
					return "", err
				}
			}
		}
		return "", nerrors.BadRequest.New("invalid verification code")
	}
}

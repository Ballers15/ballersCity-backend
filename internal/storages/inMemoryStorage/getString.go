package inMemoryStorage

import "github.com/nostressdev/nerrors"

func (s *IMStorage) GetString(accessToken string, key string) (value string, err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return "", nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	if value, ok := s.UserSettings[userEmail].StringMap[key]; !ok {
		return "", nerrors.NotFound.New("no such key in user's string map")
	} else {
		return value, nil
	}
}

package inMemoryStorage

import "github.com/nostressdev/nerrors"

func (s *IMStorage) SetString(accessToken string, key string, value string) (err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	s.UserSettings[userEmail].StringMap[key] = value
	return nil
}

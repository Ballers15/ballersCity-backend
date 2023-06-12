package inMemoryStorage

import "github.com/nostressdev/nerrors"

func (s *IMStorage) SetFloat(accessToken string, key string, value float64) (err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	s.UserSettings[userEmail].FloatMap[key] = value
	return nil
}

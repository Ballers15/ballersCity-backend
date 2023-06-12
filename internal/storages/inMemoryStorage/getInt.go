package inMemoryStorage

import "github.com/nostressdev/nerrors"

func (s *IMStorage) GetInt(accessToken string, key string) (value int64, err error) {
	if _, ok := s.UserAccessTokens[accessToken]; !ok {
		return 0, nerrors.BadRequest.New("invalid access token")
	}
	userEmail := s.UserAccessTokens[accessToken]
	if value, ok := s.UserSettings[userEmail].IntMap[key]; !ok {
		return 0, nerrors.NotFound.New("no such key in user's int map")
	} else {
		return value, nil
	}
}

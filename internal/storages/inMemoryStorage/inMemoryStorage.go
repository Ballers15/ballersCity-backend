package inMemoryStorage

import "github.com/nostressdev/tvx/pb"

type IMStorage struct {
	Products               []*pb.Product
	UserConsumedProducts   map[string][]*pb.Product //email-[]Product
	UserSettings           map[string]*UserSettings //email-UserSettings
	UserAccessTokens       map[string]string        //token-email
	EmailVerificationCodes map[string]string        //email-verification code
}

type UserSettings struct {
	StringMap map[string]string
	IntMap    map[string]int64
	FloatMap  map[string]float64
}

func NewIMStorage() IMStorage {
	return IMStorage{
		Products:               make([]*pb.Product, 0),
		UserConsumedProducts:   make(map[string][]*pb.Product),
		UserSettings:           make(map[string]*UserSettings),
		UserAccessTokens:       make(map[string]string),
		EmailVerificationCodes: make(map[string]string),
	}
}

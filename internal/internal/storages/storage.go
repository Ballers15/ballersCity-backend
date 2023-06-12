package storages

import "github.com/nostressdev/tvx/pb"

type TvxStorage interface {
	Login(email string) (accessToken string, err error)

	SetString(accessToken string, key string, value string) (err error)
	GetString(accessToken string, key string) (value string, err error)
	SetInt(accessToken string, key string, value int64) (err error)
	GetInt(accessToken string, key string) (value int64, err error)
	SetFloat(accessToken string, key string, value float64) (err error)
	GetFloat(accessToken string, key string) (value float64, err error)

	GetAvailableProductList() (products []*pb.Product)
	ConsumeProduct(accessToken string, id string) (err error)
	GetConsumedProductList(accessToken string) (products []*pb.Product, err error)
}

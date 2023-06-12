package postgresStorage

import (
	"github.com/nostressdev/tvx/pb"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) (*PostgresStorage, error) {
	err := db.AutoMigrate(&UserAccessTokensModel{}, &UserSettingsStringModel{}, &UserSettingsIntModel{}, &UserSettingsFloatModel{}, &ProductModel{}, &UserConsumedProductModel{})
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

type ProductModel struct {
	gorm.Model
	pb.Product
	PbId string
}

type UserAccessTokensModel struct {
	gorm.Model
	AccessToken string
	Email       string
}

type UserSettingsStringModel struct {
	gorm.Model
	Email string
	Key   string
	Value string
}

type UserSettingsIntModel struct {
	gorm.Model
	Email string
	Key   string
	Value int64
}

type UserSettingsFloatModel struct {
	gorm.Model
	Email string
	Key   string
	Value float64
}

type UserConsumedProductModel struct {
	gorm.Model
	pb.Product
	Email string
	PbId  string
}

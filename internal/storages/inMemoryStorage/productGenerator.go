package inMemoryStorage

import (
	"crypto/rand"
	"fmt"
	"github.com/nostressdev/tvx/internal/storages"
	"github.com/nostressdev/tvx/pb"
	"math/big"
)

func (s *IMStorage) GenerateProducts() error {
	productIds := make([]string, 20)
	for i := 0; i < 20; i++ {
		id, err := storages.GenerateOneTimePassword(32)
		if err != nil {
			return err
		}
		for j := 0; j < i; j++ {
			if productIds[j] == id {
				id, err = storages.GenerateOneTimePassword(32)
				if err != nil {
					return err
				}
				j = 0
			}
		}
		productIds[i] = id
	}
	for ind, id := range productIds {
		price, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			return err
		}
		product := &pb.Product{
			Id:           id,
			PaymentType:  pb.PaymentType((ind % 6) + 1),
			Price:        float64(price.Int64()),
			IsConsumable: ind%2 == 0,
		}
		s.Products = append(s.Products, product)
	}
	fmt.Printf("Products generated successfully!\n")
	return nil
}

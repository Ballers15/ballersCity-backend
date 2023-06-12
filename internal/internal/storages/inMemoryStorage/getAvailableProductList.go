package inMemoryStorage

import "github.com/nostressdev/tvx/pb"

func (s *IMStorage) GetAvailableProductList() (products []*pb.Product) {
	availableProducts := []*pb.Product{}
	for _, pr := range s.Products {
		if pr.IsConsumable {
			availableProducts = append(availableProducts, pr)
		}
	}
	return availableProducts
}

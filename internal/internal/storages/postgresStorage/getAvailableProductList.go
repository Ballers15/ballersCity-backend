package postgresStorage

import "github.com/nostressdev/tvx/pb"

func (s *PostgresStorage) GetAvailableProductList() (products []*pb.Product) {
	availableProductsModels := []*ProductModel{}
	s.db.Where("is_consumable = ?", "true").Find(&availableProductsModels)
	availableProducts := make([]*pb.Product, len(availableProductsModels))
	for i, product := range availableProductsModels {
		availableProducts[i] = &pb.Product{
			Id:           product.PbId,
			PaymentType:  product.PaymentType,
			Price:        product.Price,
			IsConsumable: product.IsConsumable,
		}
	}
	return availableProducts
}

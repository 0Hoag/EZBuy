package products

import (
	"context"
	"product-service/modules/item/entity"
)

type CreateProductStorage interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
}

type createProductBiz struct {
	store CreateProductStorage
}

func NewCreateProduct(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProductBiz(ctx context.Context, product *entity.Product) error {
	if err := biz.store.CreateProduct(ctx, product); err != nil {
		return entity.CantCreateProduct(err)
	}
	return nil
}

package products

import (
	"context"
	"product-service/dto"
	"product-service/modules/item/entity"
)

type UpdateProductStorage interface {
	GetProductById(ctx context.Context, cond map[string]interface{}) (error, *entity.Product)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.ProductUpdateRequest) error
}

type updateProductBiz struct {
	store UpdateProductStorage
}

func NewUpdateCategory(store UpdateProductStorage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

func (biz *updateProductBiz) UpdateProductBiz(ctx context.Context, id string, dataUpdate *dto.ProductUpdateRequest) error {
	err, _ := biz.store.GetProductById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return entity.CantGetCategory(err)
	}

	if err1 := biz.store.UpdateProduct(ctx, map[string]interface{}{"id": id}, dataUpdate); err1 != nil {
		return entity.CantUpdateCategory(err1)
	}

	return nil
}

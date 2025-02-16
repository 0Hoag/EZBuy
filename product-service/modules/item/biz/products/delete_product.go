package products

import (
	"context"
	"product-service/modules/item/entity"
)

type DeleteProductStorage interface {
	DeleteProduct(ctx context.Context, cond map[string]interface{}) error
}

type deleteProductBiz struct {
	store DeleteProductStorage
}

func NewDeleteCategory(store DeleteProductStorage) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteCategoryBiz(ctx context.Context, id string) error {
	if err := biz.store.DeleteProduct(ctx, map[string]interface{}{"id": id}); err != nil {
		return entity.CantDeleteCategory(err)
	}

	return nil
}

package categories

import (
	"context"
	"product-service/modules/item/entity"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, category *entity.Category) error
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategory(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{
		store: store,
	}
}

func (biz *createCategoryBiz) CreateCategoryBiz(ctx context.Context, data *entity.Category) error {
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return entity.CantCreateCategory(err)
	}
	return nil
}

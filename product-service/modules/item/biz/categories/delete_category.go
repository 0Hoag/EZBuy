package categories

import (
	"context"
	"product-service/modules/item/entity"
)

type DeleteCategoryStorage interface {
	DeleteCategory(ctx context.Context, cond map[string]interface{}) error
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategory(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategoryBiz(ctx context.Context, id string) error {
	if err := biz.store.DeleteCategory(ctx, map[string]interface{}{"user_id": id}); err != nil {
		return entity.CantDeleteCategory(err)
	}

	return nil
}

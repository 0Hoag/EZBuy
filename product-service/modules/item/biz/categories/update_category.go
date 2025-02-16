package categories

import (
	"context"
	"product-service/dto"
	"product-service/modules/item/entity"
)

type UpdateCategoryStorage interface {
	GetCategoryById(ctx context.Context, cond map[string]interface{}) (error, *entity.Category)
	UpdateCategory(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.CategoryUpdateRequest) error
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategory(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategoryBiz(ctx context.Context, id string, dataUpdate *dto.CategoryUpdateRequest) error {
	err, _ := biz.store.GetCategoryById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return entity.CantGetCategory(err)
	}

	if err1 := biz.store.UpdateCategory(ctx, map[string]interface{}{"id": id}, dataUpdate); err1 != nil {
		return entity.CantUpdateCategory(err1)
	}

	return nil
}

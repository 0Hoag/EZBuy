package categories

import (
	"context"
	modules "product-service/modules/item"
	"product-service/modules/item/entity"
)

type GetCategoryStorage interface {
	GetCategoryById(ctx context.Context, cond map[string]interface{}) (error, *entity.Category)
	GetAllCategory(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Category)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

func NewGetCategoryBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategoryByIdBiz(ctx context.Context, id string) (error, *entity.Category) {
	err, data := biz.store.GetCategoryById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return entity.CantGetCategory(err), nil
	}

	return nil, data
}

func (biz *getCategoryBiz) GetAllCategoryBiz(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Category) {
	err, data := biz.store.GetAllCategory(ctx, paging)
	if err != nil {
		return entity.CantGetListCategory(err), nil
	}

	return nil, data
}

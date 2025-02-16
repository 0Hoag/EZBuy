package products

import (
	"context"
	modules "product-service/modules/item"
	"product-service/modules/item/entity"
)

type GetProductStorage interface {
	GetProductById(ctx context.Context, cond map[string]interface{}) (error, *entity.Product)
	GetAllProduct(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Product)
}

type getProductBiz struct {
	store GetProductStorage
}

func NewGetCategoryBiz(store GetProductStorage) *getProductBiz {
	return &getProductBiz{store: store}
}

func (biz *getProductBiz) GetProductByIdBiz(ctx context.Context, id string) (error, *entity.Product) {
	err, data := biz.store.GetProductById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return entity.CantGetCategory(err), nil
	}

	return nil, data
}

func (biz *getProductBiz) GetAllProductBiz(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Product) {
	err, data := biz.store.GetAllProduct(ctx, paging)
	if err != nil {
		return entity.CantGetListCategory(err), nil
	}

	return nil, data
}

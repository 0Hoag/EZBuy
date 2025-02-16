package user

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

type GetUserStorage interface {
	GetUserById(ctx context.Context, cond map[string]interface{}) (*entity.User, error)
	GetUserByUserName(ctx context.Context, cond map[string]interface{}) (*entity.User, error)
	GetAllUser(ctx context.Context, paging *modules.Paging, moreKeys ...string) ([]entity.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserByIdBiz(ctx context.Context, id string) (*entity.User, error) {
	data, err := biz.store.GetUserById(ctx, map[string]interface{}{"user_id": id})

	if err != nil {
		return nil, entity.CantGetUser(err)
	}

	return data, nil
}

func (biz *getUserBiz) GetUserByUserNameBiz(ctx context.Context, username string) (*entity.User, error) {
	data, err := biz.store.GetUserByUserName(ctx, map[string]interface{}{"username": username})

	if err != nil {
		return nil, entity.ErrUserNameBlank
	}

	return data, nil
}

func (biz *getUserBiz) GetAllUser(ctx context.Context, paging *modules.Paging, moreKeys ...string) ([]entity.User, error) {
	data, err := biz.store.GetAllUser(ctx, paging)

	if err != nil {
		return nil, entity.CantGetListUser(err)
	}

	return data, nil
}

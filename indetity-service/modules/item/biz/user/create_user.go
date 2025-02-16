package user

import (
	"context"
	"indetity-service/modules/item/entity"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *entity.User) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateNewBiz(ctx context.Context, data *entity.User) error {
	if data.Username == "" {
		return entity.ErrUserNameBlank
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return entity.CantCreateUser(err)
	}
	return nil
}

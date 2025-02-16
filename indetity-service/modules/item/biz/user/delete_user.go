package user

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

type DeleteUserStorage interface {
	GetUserById(ctx context.Context, cond map[string]interface{}) (*entity.User, error)
	DeleteUserById(ctx context.Context, cond map[string]interface{}) error
}

type deleteUserBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUserBiz(ctx context.Context, id string) error {
	_, err := biz.store.GetUserById(ctx, map[string]interface{}{"user_id": id})

	if err != nil {
		if err == modules.RecordNotFound {
			return entity.CantGetUser(err)
		}

		return entity.CantDeleteUser(err)
	}

	if err1 := biz.store.DeleteUserById(ctx, map[string]interface{}{"user_id": id}); err1 != nil {
		return entity.CantDeleteUser(err1)
	}

	return nil
}

package user

import (
	"context"
	"indetity-service/dto"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

type UpdateUserStorage interface {
	GetUserById(ctx context.Context, cond map[string]interface{}) (*entity.User, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoUpdateRequest) error
}

type UpdateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *UpdateUserBiz {
	return &UpdateUserBiz{store: store}
}

func (biz *UpdateUserBiz) UpdateUserBiz(ctx context.Context, id string, dataUpdate *dto.TodoUpdateRequest) error {
	_, err := biz.store.GetUserById(ctx, map[string]interface{}{"user_id": id})

	if err != nil {
		if err == modules.RecordNotFound {
			return entity.CantGetUser(err)
		}

		return entity.CantUpdateUser(err)
	}

	if err1 := biz.store.UpdateUser(ctx, map[string]interface{}{"user_id": id}, dataUpdate); err1 != nil {
		return entity.CantUpdateUser(err1)
	}

	return nil
}

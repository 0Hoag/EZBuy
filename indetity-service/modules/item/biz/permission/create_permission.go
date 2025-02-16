package permission

import (
	"context"
	"errors"
	"indetity-service/modules/item/entity"
)

type CreatePermissionStorage interface {
	CreatePermission(ctx context.Context, permission *entity.Permission) error
}

type createPermissionBiz struct {
	store CreatePermissionStorage
}

func NewCreatePermission(store CreatePermissionStorage) *createPermissionBiz {
	return &createPermissionBiz{store: store}
}

func (biz *createPermissionBiz) CreatePermissionBiz(ctx context.Context, data *entity.Permission) error {
	if data.Name == "" {
		return errors.New("Permission name it not empty!")
	}

	err := biz.store.CreatePermission(ctx, data)
	if err != nil {
		return entity.CantCreatePermission(err)
	}
	return nil
}

package permission

import (
	"context"
	"indetity-service/dto"
	"indetity-service/modules/item/entity"
)

type UpdatePermissionStorage interface {
	GetPermission(ctx context.Context, cond map[string]interface{}) (*entity.Permission, error)
	UpdatePermission(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoPermissionUpdate) error
}

type UpdatePermission struct {
	store UpdatePermissionStorage
}

func NewUpdatePermission(store UpdatePermissionStorage) *UpdatePermission {
	return &UpdatePermission{store: store}
}

func (biz *UpdatePermission) UpdatePermissionBiz(ctx context.Context, id string, dataUpdate *dto.TodoPermissionUpdate) error {
	_, err := biz.store.GetPermission(ctx, map[string]interface{}{"name": id})

	if err != nil {
		return entity.CantGetPermission(err)
	}

	if err1 := biz.store.UpdatePermission(ctx, map[string]interface{}{"name": id}, dataUpdate); err1 != nil {
		return entity.CantUpdatePermission(err)
	}

	return nil
}

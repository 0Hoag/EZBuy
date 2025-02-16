package permission

import (
	"context"
	"indetity-service/modules/item/entity"
)

type DeletePermissionStorage interface {
	DeletePermission(ctx context.Context, cond map[string]interface{}) error
}

type deletePermission struct {
	store DeletePermissionStorage
}

func NewDeletePermission(store DeletePermissionStorage) *deletePermission {
	return &deletePermission{store: store}
}

func (biz *deletePermission) DeletePermissionBiz(ctx context.Context, id string) error {
	if err := biz.store.DeletePermission(ctx, map[string]interface{}{"name": id}); err != nil {
		return entity.CantDeletePermission(err)
	}

	return nil
}

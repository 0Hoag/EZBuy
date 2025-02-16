package permission

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

type GetPermissionStorage interface {
	GetPermission(ctx context.Context, cond map[string]interface{}) (*entity.Permission, error)
	GetAllPermission(ctx context.Context, paging *modules.Paging) ([]entity.Permission, error)
}

type getPermission struct {
	store GetPermissionStorage
}

func NewGetPermission(store GetPermissionStorage) *getPermission {
	return &getPermission{store: store}
}

func (biz *getPermission) GetPermissionBiz(ctx context.Context, name string) (*entity.Permission, error) {
	data, err := biz.store.GetPermission(ctx, map[string]interface{}{"name": name})
	if err != nil {
		return nil, entity.CantGetPermission(err)
	}

	return data, nil
}

func (biz *getPermission) GetAllPermissionBiz(ctx context.Context, paging *modules.Paging) ([]entity.Permission, error) {
	data, err := biz.store.GetAllPermission(ctx, paging)
	if err != nil {
		return nil, entity.CantGetListPermission(err)
	}

	return data, nil
}

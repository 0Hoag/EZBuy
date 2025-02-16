package role

import (
	"context"
	"indetity-service/modules/item/entity"
)

type DeleteRoleStorage interface {
	DeleteRole(ctx context.Context, cond map[string]interface{}) error
}

type deleteRoleBiz struct {
	store DeleteRoleStorage
}

func NewDeleteRoleBiz(store DeleteRoleStorage) *deleteRoleBiz {
	return &deleteRoleBiz{store: store}
}

func (biz *deleteRoleBiz) DeleteRoleBiz(ctx context.Context, id string) error {
	if err1 := biz.store.DeleteRole(ctx, map[string]interface{}{"name": id}); err1 != nil {
		return entity.CantDeleteRole(err1)
	}

	return nil
}

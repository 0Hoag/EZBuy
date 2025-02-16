package role

import (
	"context"
	"indetity-service/modules/item/entity"
)

type CreateRoleStorage interface {
	CreateRole(ctx context.Context, role *entity.Role) error
}

type createRoleBiz struct {
	store CreateRoleStorage
}

func NewCreateRole(store CreateRoleStorage) *createRoleBiz {
	return &createRoleBiz{store: store}
}

func (biz *createRoleBiz) CreateRoleBiz(ctx context.Context, role *entity.Role) error {
	err := biz.store.CreateRole(ctx, role)
	if err != nil {
		return entity.CantCreateRole(err)
	}

	return nil
}

package role

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

type GetRoleStorage interface {
	GetRoleById(ctx context.Context, cond map[string]interface{}) (*entity.Role, error)
	GetAllRole(ctx context.Context, paging *modules.Paging) ([]entity.Role, error)
}

type getRoleBiz struct {
	store GetRoleStorage
}

func NewGetRole(store GetRoleStorage) *getRoleBiz {
	return &getRoleBiz{store: store}
}

func (biz *getRoleBiz) GetRoleBiz(ctx context.Context, name string) (*entity.Role, error) {
	data, err := biz.store.GetRoleById(ctx, map[string]interface{}{"name": name})
	if err != nil {
		return nil, entity.CantGetRole(err)
	}

	return data, nil
}

func (biz *getRoleBiz) GetAllRoleBiz(ctx context.Context, paging *modules.Paging) ([]entity.Role, error) {
	data, err := biz.store.GetAllRole(ctx, paging)
	if err != nil {
		return nil, entity.CantGetListRole(err)
	}

	return data, nil
}

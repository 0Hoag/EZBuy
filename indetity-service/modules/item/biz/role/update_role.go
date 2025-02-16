package role

import (
	"context"
	"errors"
	"indetity-service/dto"
	"indetity-service/modules/item/entity"
)

type UpdateRoleStorage interface {
	GetRoleById(ctx context.Context, cond map[string]interface{}) (*entity.Role, error)
	UpdateRole(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoRoleUpdate) error
}

type UpdateRoleBiz struct {
	store UpdateRoleStorage
}

func NewUpdateRoleBiz(store UpdateRoleStorage) *UpdateRoleBiz {
	return &UpdateRoleBiz{store: store}
}

func (biz *UpdateRoleBiz) UpdateRoleBiz(ctx context.Context, id string, dataUpdate *dto.TodoRoleUpdate) error {
	_, err := biz.store.GetRoleById(ctx, map[string]interface{}{"name": id})
	if err != nil {
		return entity.CantGetRole(err)
	}

	if dataUpdate.Name != nil && *dataUpdate.Name == "" {
		return errors.New("role name cannot be empty")
	}

	if dataUpdate.Permission != nil && len(*dataUpdate.Permission) == 0 {
		return errors.New("permissions list cannot be empty")
	}

	return biz.store.UpdateRole(ctx, map[string]interface{}{"name": id}, dataUpdate)
}

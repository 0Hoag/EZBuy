package entity

import (
	"errors"
	"indetity-service/modules"
)

const (
	EntityRole       = "role"
	EntityPermission = "permission"
)

var (
	RoleNotFound    = modules.RecordNotFound
	CantCreateRole  = func(err error) error { return modules.ErrCannotCreateEntity(EntityRole, err) }
	CantGetRole     = func(err error) error { return modules.ErrCannotGetEntity(EntityRole, err) }
	CantGetListRole = func(err error) error { return modules.ErrCannotListEntity(EntityRole, err) }
	CantUpdateRole  = func(err error) error { return modules.ErrCannotUpdateEntity(EntityRole, err) }
	CantDeleteRole  = func(err error) error { return modules.ErrCannotDeleteEntity(EntityRole, err) }
)

var (
	PermissionNotFound    = modules.RecordNotFound
	CantCreatePermission  = func(err error) error { return modules.ErrCannotCreateEntity(EntityPermission, err) }
	CantGetPermission     = func(err error) error { return modules.ErrCannotGetEntity(EntityPermission, err) }
	CantGetListPermission = func(err error) error { return modules.ErrCannotListEntity(EntityPermission, err) }
	CantUpdatePermission  = func(err error) error { return modules.ErrCannotUpdateEntity(EntityPermission, err) }
	CantDeletePermission  = func(err error) error { return modules.ErrCannotDeleteEntity(EntityPermission, err) }
)

var (
	UnAuthentication = modules.ErrUnAuthorized(EntityRole, errors.New("UnAuthorization"))
	TokenInvalid     = modules.ErrTokenInvalid(errors.New("Token invalid!"))
)

type Role struct {
	Name        string       `json:"name" gorm:"column:name;primaryKey"`
	Description string       `json:"description" gorm:"column:description"`
	Permission  []Permission `json:"permission" gorm:"many2many:role_permissions"`
}

func (Role) TableName() string {
	return "role"
}

type Permission struct {
	Name        string `json:"name" gorm:"column:name;primaryKey"`
	Description string `json:"description" gorm:"column:description"`
}

func (Permission) TableName() string {
	return "permission"
}

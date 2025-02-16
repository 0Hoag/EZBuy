package entity

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"indetity-service/modules"
)

const (
	EntityUser = "user"
)

var (
	ErrUserNameBlank = errors.New("username it blank")
	UserNotFound     = modules.RecordNotFound
	CantCreateUser   = func(err error) error { return modules.ErrCannotCreateEntity(EntityUser, err) }
	CantGetUser      = func(err error) error { return modules.ErrCannotGetEntity(EntityUser, err) }
	CantGetListUser  = func(err error) error { return modules.ErrCannotListEntity(EntityUser, err) }
	CantUpdateUser   = func(err error) error { return modules.ErrCannotUpdateEntity(EntityUser, err) }
	CantDeleteUser   = func(err error) error { return modules.ErrCannotDeleteEntity(EntityUser, err) }
	ErrorDB          = func(err error) error { return modules.ErrDB(err) }
)

type User struct {
	UserId      string `json:"user_id" gorm:"column:user_id;primaryKey"`
	Username    string `json:"username" gorm:"column:username;unique;not null"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Password    string `json:"password" gorm:"column:password"`
	Email       string `json:"email" gorm:"column:email"`
	EmailVerify bool   `json:"email_verify" gorm:"column:email_verified"`
	Role        []Role `json:"role" gorm:"many2many:user_roles"`
}

type UserRoles struct {
	UserId   string `json:"user_user_id" gorm:"column:user_user_id;primaryKey"`
	RoleName string `json:"role_name" gorm:"column:role_name;primaryKey"`
}

type RolePermissions struct {
	RoleName       string `json:"role_name" gorm:"column:role_name;primaryKey"`
	PermissionName string `json:"permission_name" gorm:"column:permission_name;primaryKey"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}

func (RolePermissions) TableName() string {
	return "role_permissions"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserId == "" {
		u.UserId = uuid.New().String()
	}
	return
}

func (User) TableName() string {
	return "user"
}

package dto

import (
	"indetity-service/modules/item/entity"
)

type UserRequest struct {
	Username  string `json:"username" gorm:"column:username;unique;not null"`
	FirstName string `json:"first_name,omitempty" gorm:"column:first_name"`
	LastName  string `json:"last_name,omitempty" gorm:"column:last_name"`
	Password  string `json:"password" gorm:"column:password"`
	Email     string `json:"email" gorm:"column:email"`
}

type UserResponse struct {
	UserId      string        `json:"user_id"`
	Username    string        `json:"username"`
	FirstName   string        `json:"first_name,omitempty"`
	LastName    string        `json:"last_name,omitempty"`
	Email       string        `json:"email"`
	EmailVerify bool          `json:"email_verify"`
	Role        []entity.Role `json:"role"`
}

type TodoUpdateRequest struct {
	FirstName *string `json:"first_name" gorm:"column:first_name"`
	LastName  *string `json:"last_name" gorm:"column:last_name"`
	Password  *string `json:"password" gorm:"column:password"`
	Email     *string `json:"email" gorm:"column:email"`
}

func (TodoUpdateRequest) TableName() string {
	return "user"
}

func (req *UserRequest) ToUser() *entity.User { // mapstruct
	return &entity.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Email:     req.Email,
	}
}

func ToUserResponse(user *entity.User) *UserResponse { // mapstruct
	return &UserResponse{
		UserId:      user.UserId,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		EmailVerify: user.EmailVerify,
		Role:        user.Role,
	}
}

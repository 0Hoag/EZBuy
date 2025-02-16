package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	modules "product-service/modules/item"
)

const (
	DatabaseName       = "product-service"
	CollectionCategory = "category"
)

var (
	CategoryNotFound    = modules.RecordNotFound
	CantCreateCategory  = func(err error) error { return modules.ErrCannotCreateEntity(CollectionCategory, err) }
	CantGetCategory     = func(err error) error { return modules.ErrCannotGetEntity(CollectionCategory, err) }
	CantGetListCategory = func(err error) error { return modules.ErrCannotListEntity(CollectionCategory, err) }
	CantUpdateCategory  = func(err error) error { return modules.ErrCannotUpdateEntity(CollectionCategory, err) }
	CantDeleteCategory  = func(err error) error { return modules.ErrCannotDeleteEntity(CollectionCategory, err) }
	ErrorDB             = func(err error) error { return modules.ErrDB(err) }
	ErrorResourceLeak   = func(err error) error { return modules.ErrInternal(err) }
)

type Category struct {
	Id               string `json:"id" gorm:"column:id;primaryKey"`
	Name             string `json:"name" gorm:"column:name;unique;not null"`
	Description      string `json:"description" gorm:"column:description"`
	ParentCategoryId string `json:"parent_category_id" gorm:"parent_category_id"`
}

func (u *Category) BeforeCreate(tx *mongo.Client) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}

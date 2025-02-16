package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	modules "product-service/modules/item"
)

const (
	CollectionProduct = "product"
)

var (
	ProductNotFound    = modules.RecordNotFound
	CantCreateProduct  = func(err error) error { return modules.ErrCannotCreateEntity(CollectionProduct, err) }
	CantGetProduct     = func(err error) error { return modules.ErrCannotGetEntity(CollectionProduct, err) }
	CantGetListProduct = func(err error) error { return modules.ErrCannotListEntity(CollectionProduct, err) }
	CantUpdateProduct  = func(err error) error { return modules.ErrCannotUpdateEntity(CollectionProduct, err) }
	CantDeleteProduct  = func(err error) error { return modules.ErrCannotDeleteEntity(CollectionProduct, err) }
)

type Product struct {
	Id            string   `json:"id" gorm:"column:id;primaryKey"`
	Author        string   `json:"author" gorm:"column:author"`
	Name          string   `json:"name" gorm:"column:name;unique;not null"`
	Address       string   `json:"address" gorm:"column:address"`
	ListedPrice   float32  `json:"listed_price" gorm:"column:listed_price"`
	Price         float32  `json:"price" gorm:"column:price"`
	Quantity      int      `json:"quantity" gorm:"column:quantity"`
	StockQuantity int      `json:"stock_quantity" gorm:"column:stock_quantity"`
	Description   string   `json:"description" gorm:"column:description"`
	Image         []string `json:"image" gorm:"column:image"`
	CategoryId    []string `json:"category_id" gorm:"column:category_id"`
}

func (u *Product) BeforeCreate(tx *mongo.Client) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}

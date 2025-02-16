package entity

import (
	modules "cartitem-service/modules/item"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CollectionCartItem = "cart_item"
)

var (
	CartItemNotFound   = modules.RecordNotFound
	CantCreateCartIte  = func(err error) error { return modules.ErrCannotCreateEntity(CollectionCartItem, err) }
	CantGetCartIte     = func(err error) error { return modules.ErrCannotGetEntity(CollectionCartItem, err) }
	CantGetListCartIte = func(err error) error { return modules.ErrCannotListEntity(CollectionCartItem, err) }
	CantUpdateCartIte  = func(err error) error { return modules.ErrCannotUpdateEntity(CollectionCartItem, err) }
	CantDeleteCartIte  = func(err error) error { return modules.ErrCannotDeleteEntity(CollectionCartItem, err) }
)

type CartItem struct {
	CartItemId string `json:"cart_item_id" gorm:"column:cart_item_id;primaryKey"`
	UserId     string `json:"user_id" gorm:"column:user_id"`
	Quantity   int    `json:"quantity" gorm:"column:quantity"`
	ProductId  string `json:"product_id" gorm:"column:product_id"`
}

func (u *CartItem) BeforeCreate(tx *mongo.Client) (err error) {
	if u.CartItemId == "" {
		u.CartItemId = uuid.New().String()
	}
	return
}

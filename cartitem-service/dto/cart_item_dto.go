package dto

import "cartitem-service/modules/item/entity"

type CartItemRequest struct {
	userId    string `json:"user_id"`
	quantity  int    `json:"quantity"`
	productId string `json:"product_id"`
}

type CartItemResponse struct {
	cartItemId string `json:"cart_item_id"`
	userId     string `json:"user_id"`
	quantity   int    `json:"quantity"`
	productId  string `json:"product_id"`
}

type CartItemUpdateRequest struct {
	userId    string `json:"user_id"`
	quantity  int    `json:"quantity"`
	productId string `json:"product_id"`
}

func (req *CartItemRequest) ToCartItem() *entity.CartItem { // mapstruct
	return &entity.CartItem{
		UserId:    req.userId,
		Quantity:  req.quantity,
		ProductId: req.productId,
	}
}

func ToCartItemResponse(cart *entity.CartItem) *CartItemResponse { // mapstruct
	return &CartItemResponse{
		cartItemId: cart.CartItemId,
		userId:     cart.UserId,
		quantity:   cart.Quantity,
		productId:  cart.ProductId,
	}
}

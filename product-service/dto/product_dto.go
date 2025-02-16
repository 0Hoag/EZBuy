package dto

import "product-service/modules/item/entity"

type ProductRequest struct {
	Author        string   `json:"author"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	ListedPrice   float32  `json:"listed_price"`
	Price         float32  `json:"price"`
	Quantity      int      `json:"quantity"`
	StockQuantity int      `json:"stock_quantity"`
	Description   string   `json:"description"`
	Image         []string `json:"image"`
	CategoryId    []string `json:"category_id"`
}

type ProductResponse struct {
	Id            string   `json:"id"`
	Author        string   `json:"author"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	ListedPrice   float32  `json:"listed_price"`
	Price         float32  `json:"price"`
	Quantity      int      `json:"quantity"`
	StockQuantity int      `json:"stock_quantity"`
	Description   string   `json:"description"`
	Image         []string `json:"image"`
	CategoryId    []string `json:"category_id"`
}

type ProductUpdateRequest struct {
	Author        string   `json:"author"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	ListedPrice   float32  `json:"listed_price"`
	Price         float32  `json:"price"`
	Quantity      int      `json:"quantity"`
	StockQuantity int      `json:"stock_quantity"`
	Description   string   `json:"description"`
	Image         []string `json:"image"`
	CategoryId    []string `json:"category_id"`
}

func (req *ProductRequest) ToProduct() *entity.Product {
	return &entity.Product{
		Author:        req.Author,
		Name:          req.Name,
		Address:       req.Address,
		ListedPrice:   req.ListedPrice,
		Price:         req.Price,
		Quantity:      req.Quantity,
		StockQuantity: req.StockQuantity,
		Description:   req.Description,
		Image:         req.Image,
		CategoryId:    req.CategoryId,
	}
}

func ToProductResponse(pro *entity.Product) *ProductResponse {
	return &ProductResponse{
		Author:        pro.Author,
		Name:          pro.Name,
		Address:       pro.Address,
		ListedPrice:   pro.ListedPrice,
		Price:         pro.Price,
		Quantity:      pro.Quantity,
		StockQuantity: pro.StockQuantity,
		Description:   pro.Description,
		Image:         pro.Image,
		CategoryId:    pro.CategoryId,
	}
}

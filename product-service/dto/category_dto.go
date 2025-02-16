package dto

import "product-service/modules/item/entity"

type CategoryRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	ParentCategoryId string `json:"parent_category_id"`
}

type CategoryUpdateRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	ParentCategoryId string `json:"parent_category_id"`
}

type CategoryResponse struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ParentCategoryId string `json:"parent_category_id"`
}

func (req *CategoryRequest) ToCategory() *entity.Category {
	return &entity.Category{
		Name:             req.Name,
		Description:      req.Description,
		ParentCategoryId: req.ParentCategoryId,
	}
}

func ToCategoryResponse(cate *entity.Category) *CategoryResponse { // mapstruct
	return &CategoryResponse{
		Id:               cate.Id,
		Name:             cate.Name,
		Description:      cate.Description,
		ParentCategoryId: cate.ParentCategoryId,
	}
}

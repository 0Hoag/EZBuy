package storage

import (
	"context"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) CreateCategory(ctx context.Context, dataCreate *entity.Category) error {
	collection := sql.db.Database("product-service").Collection("category")

	_, err := collection.InsertOne(ctx, dataCreate)

	if err != nil {
		return entity.CantCreateCategory(err)
	}

	return nil
}

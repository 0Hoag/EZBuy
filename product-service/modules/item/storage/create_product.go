package storage

import (
	"context"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) CreateProduct(ctx context.Context, product *entity.Product) error {
	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionProduct)

	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		return entity.CantCreateProduct(err)
	}

	return nil
}

package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) DeleteProduct(ctx context.Context, cond map[string]interface{}) error {
	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionProduct)
	filter := bson.M{"id": cond["id"]}

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return entity.CantDeleteProduct(err)
	}

	return nil
}

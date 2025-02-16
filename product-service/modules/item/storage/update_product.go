package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"product-service/dto"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) UpdateCategory(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.CategoryUpdateRequest) error {
	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionProduct)

	filter := bson.M{
		"id": cond["id"],
	}

	opts := options.Update().SetUpsert(false)

	_, err := collection.UpdateOne(ctx, filter, dataUpdate, opts)
	if err != nil {
		return entity.CantUpdateCategory(err)
	}

	return nil
}

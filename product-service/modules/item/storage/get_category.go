package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	modules "product-service/modules/item"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) GetAllCategory(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Category) {
	data := make([]*entity.Category, 0)

	paging.Process()

	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionCategory)

	total, err := collection.CountDocuments(ctx, bson.D{})

	paging.TotalPage = total

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return entity.CantGetListCategory(err), nil
	}

	defer cursor.Close(ctx)

	if err1 := cursor.All(ctx, &data); err1 != nil {
		return entity.CantGetListCategory(err), nil
	}

	return nil, data
}

func (sql *SqlStorage) GetCategoryById(ctx context.Context, cond map[string]interface{}) (error, *entity.Category) {
	var data *entity.Category

	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionCategory)
	filter := bson.M{"id": cond["id"]}

	if err := collection.FindOne(ctx, filter).Decode(&data); err != nil {
		return entity.CantGetCategory(err), nil
	}

	return nil, data
}

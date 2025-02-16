package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	modules "product-service/modules/item"
	"product-service/modules/item/entity"
)

func (sql *SqlStorage) GetAllProduct(ctx context.Context, paging *modules.Paging, moreKeys ...string) (error, []*entity.Product) {
	data := make([]*entity.Product, 0)

	paging.Process()

	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionProduct)

	total, err := collection.CountDocuments(ctx, bson.D{})

	paging.TotalPage = total

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return entity.CantGetListProduct(err), nil
	}

	defer cursor.Close(ctx)

	if err1 := cursor.All(ctx, &data); err1 != nil {
		return entity.CantGetListProduct(err), nil
	}

	return nil, data
}

func (sql *SqlStorage) GetProductById(ctx context.Context, cond map[string]interface{}) (error, *entity.Product) {
	var data *entity.Product

	collection := sql.db.Database(entity.DatabaseName).Collection(entity.CollectionProduct)
	filter := bson.M{"id": cond["id"]}

	if err := collection.FindOne(ctx, filter).Decode(&data); err != nil {
		return entity.CantGetProduct(err), nil
	}

	return nil, data
}

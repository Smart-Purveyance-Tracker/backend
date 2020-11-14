package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
)

type Product interface {
	Insert(user entity.User) (entity.User, error)
	Find(id string) (entity.User, error)
}

type ProductMongoDB struct {
	collection *mongo.Collection
}

func NewProductMongoDB(collection *mongo.Collection) *ProductMongoDB {
	return &ProductMongoDB{collection: collection}
}

func (p *ProductMongoDB) Insert(product entity.Product) (entity.Product, error) {
	id, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return entity.Product{}, err
	}
	product.ID = id.InsertedID.(primitive.ObjectID).Hex()
	return product, nil
}

func (p *ProductMongoDB) Find(id string) (entity.Product, error) {
	var product entity.Product
	err := p.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	return product, err
}

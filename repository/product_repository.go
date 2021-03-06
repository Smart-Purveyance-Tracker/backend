package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
)

type Product interface {
	Insert(product entity.Product) (entity.Product, error)
	Find(id string) (entity.Product, error)
	Delete(id string) error
	Update(product entity.Product) (entity.Product, error)
	List(args ProductListArgs) ([]entity.Product, error)
}

type ProductListArgs struct {
	UserID *string
	Date   *time.Time
}

type ProductMongoDB struct {
	collection *mongo.Collection
}

func NewProductMongoDB(client *mongo.Client) *ProductMongoDB {
	collection := client.Database("purveyance").Collection("products")
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
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Product{}, err
	}
	err = p.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	return product, err
}

func (p *ProductMongoDB) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (p *ProductMongoDB) Update(product entity.Product) (entity.Product, error) {
	update := bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"type":     product.Type,
			"boughtAt": product.BoughtAt,
			"inStock":  product.InStock,
		},
	}
	objID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return entity.Product{}, err
	}
	_, err = p.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return product, err
}

func (p *ProductMongoDB) List(args ProductListArgs) ([]entity.Product, error) {
	filter := bson.M{}
	if args.UserID != nil {
		filter["userId"] = *args.UserID
	}
	if args.Date != nil {
		date := *args.Date
		rounded := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		filter["boughtAt"] = bson.M{
			"$gte": rounded,
			"$lt":  rounded.Add(time.Hour * 24),
		}
	}

	cur, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var products []entity.Product
	for cur.Next(context.TODO()) {
		var p entity.Product
		err = cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

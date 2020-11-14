package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
)

type User interface {
	Insert(user entity.User) (entity.User, error)
	Find(id string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type UserMongoDB struct {
	collection *mongo.Collection
}

func NewUserMongoDB(client *mongo.Client) *UserMongoDB {
	collection := client.Database("purveyance").Collection("users")
	mod := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		}, Options: options.Index().SetUnique(true),
	}
	_, _ = collection.Indexes().CreateOne(context.TODO(), mod)
	return &UserMongoDB{collection: collection}
}

func (u *UserMongoDB) Insert(user entity.User) (entity.User, error) {
	id, err := u.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return entity.User{}, err
	}
	user.ID = id.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func (u *UserMongoDB) Find(id string) (entity.User, error) {
	var user entity.User
	err := u.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (u *UserMongoDB) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := u.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}

package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
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

type UserInMem struct {
	mx           sync.RWMutex
	userIDToInfo map[string]entity.User
}

func NewUserInMem() *UserInMem {
	return &UserInMem{
		userIDToInfo: make(map[string]entity.User),
	}
}

func (m *UserInMem) Insert(user entity.User) (entity.User, error) {
	id := uuid.New().String()
	user.ID = id
	_, err := m.FindByEmail(user.Email)
	if err == nil {
		return entity.User{}, errors.New("already exists")
	}
	m.mx.Lock()
	m.userIDToInfo[id] = user
	m.mx.Unlock()
	return user, nil
}

func (m *UserInMem) Find(id string) (entity.User, error) {
	m.mx.RLock()
	user := m.userIDToInfo[id]
	m.mx.RUnlock()
	return user, nil
}

func (m *UserInMem) FindByEmail(email string) (entity.User, error) {
	m.mx.RLock()
	for _, user := range m.userIDToInfo {
		if user.Email == email {
			return user, nil
		}
	}
	m.mx.RUnlock()
	return entity.User{}, errors.New("not found")
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

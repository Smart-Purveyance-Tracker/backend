package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID       string    `json:"id" bson:"-"`
	Name     string    `json:"name" bson:"name"`
	Type     string    `json:"type" bson:"type"`
	BoughtAt time.Time `json:"boughtAt" bson:"boughtAt"`
	UserID   string    `json:"userId" bson:"userId"`
	InStock  bool      `json:"inStock" bson:"inStock"`
}

func (p *Product) UnmarshalBSON(data []byte) error {
	type tmpProduct Product
	bProduct := &struct {
		ID      primitive.ObjectID `json:"id" bson:"_id, inline"`
		TmpUser *tmpProduct        `bson:",inline"`
	}{
		TmpUser: (*tmpProduct)(p),
	}
	if err := bson.Unmarshal(data, &bProduct); err != nil {
		return err
	}
	p.ID = bProduct.ID.Hex()
	return nil
}

package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID string `json:"id" bson:"-"`
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

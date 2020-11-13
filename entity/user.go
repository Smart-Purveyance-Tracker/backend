package entity

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       string `json:"id" bson:"-"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) UnmarshalBSON(data []byte) error {
	type tmpUser User
	bUser := &struct {
		ID      primitive.ObjectID `json:"id" bson:"_id, inline"`
		TmpUser *tmpUser           `bson:",inline"`
	}{
		TmpUser: (*tmpUser)(u),
	}
	if err := bson.Unmarshal(data, &bUser); err != nil {
		return err
	}
	u.ID = bUser.ID.Hex()
	return nil
}

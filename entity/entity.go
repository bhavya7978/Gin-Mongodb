package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	City    string `json:"city"  bson:"city"`
	State   string `json:"state"  bson:"state"`
	Pincode int64  `json:"pincode"  bson:"pincode"`
}

type User struct {
	ID      primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	Name    string             `json:"name"  bson:"name"`
	Age     int64              `json:"age"  bson:"age"`
	Address Address            `json:"address"  bson:"address"`
}

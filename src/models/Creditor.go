package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Creditor struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" csv:"Name" bson:"name"`
	Address  string             `json:"address" csv:"Address" bson:"address"`
	Postcode string             `json:"postcode" csv:"Postcode" bson:"postcode"`
	Phone    string             `json:"phone" csv:"Phone" bson:"phone"`
	Credit   float64            `json:"credit" csv:"Credit Limit" bson:"credit"`
	Birthday string             `json:"birthday" csv:"Birthday" bson:"birthday"`
}

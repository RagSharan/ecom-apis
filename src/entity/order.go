package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id      primitive.ObjectID `bson:"_id"`
	CustId  string             `bson:"custId"`
	Address string             `bson:"address"`
	Items   []Product          `bson:"items"`
	Status  string             `bson:"status"`
}

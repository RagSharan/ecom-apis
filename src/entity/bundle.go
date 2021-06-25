package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bundle struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Desc       string             `bson:"desc"`
	ProductIds []string           `bson:"productIds"`
}

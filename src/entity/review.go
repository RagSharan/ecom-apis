package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id        primitive.ObjectID `bson:"_id"`
	ProductId string             `bson:"productId"`
	Rating    int                `bson:"rating"`
	Comment   string             `bson:"comment"`
}

package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductId primitive.ObjectID `bson:"_id"`
	ProdName  string             `bson:"productname"`
	Decr      string             `bson:"descr,omitempty"`
	PriceList []Price            `bson:"priceList,omitempty"`
}

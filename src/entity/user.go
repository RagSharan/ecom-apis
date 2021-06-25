package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	UserName string             `bson:"username"`
	Email    string             `bson:"email"`
	Phone    string             `bson:"phone"`
	Address  string             `bson:"address"`
}

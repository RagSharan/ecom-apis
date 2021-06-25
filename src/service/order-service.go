package service

import (
	"context"
	"log"

	"github.com/ragsharan/ecom-apis/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderService struct{}
type IOrderService interface {
	CreateOrder(order entity.Order) (*mongo.InsertOneResult, error)
	GetOrder(order entity.Order) (entity.Order, error)
	GetOrderList(order entity.Order) ([]entity.Order, error)
	UpdateOrder()
	RemoveOrder(order entity.Order) (*mongo.DeleteResult, error)
}

const OCollection string = "order"

func InstOrderService() IOrderService {
	return &orderService{}
}

func (*orderService) CreateOrder(order entity.Order) (*mongo.InsertOneResult, error) {
	result, err := repoMongo.Create(OCollection, order)
	return result, err
}
func (*orderService) GetOrder(order entity.Order) (entity.Order, error) {
	result := repoMongo.FindOne(OCollection, order)
	err := result.Decode(&order)
	return order, err
}
func (*orderService) GetOrderList(order entity.Order) ([]entity.Order, error) {
	var orderList []entity.Order
	result, err := repoMongo.FindList(OCollection, order)
	if err != nil {
		log.Println(err)
	}
	result.All(context.TODO(), &orderList)
	return orderList, err

}

func (*orderService) RemoveOrder(order entity.Order) (*mongo.DeleteResult, error) {
	result, err := repoMongo.DeleteDocument(OCollection, order)
	return result, err
}

func (*orderService) UpdateOrder() {}

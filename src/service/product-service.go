package service

import (
	"context"
	"log"

	"github.com/ragsharan/ecom-apis/entity"
	"github.com/ragsharan/ecom-apis/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type productService struct{}

type IProductService interface {
	AddProduct(product entity.Product) (*mongo.InsertOneResult, error)
	AddProductList(product []interface{}) (*mongo.InsertManyResult, error)
	GetProduct(product entity.Product) (entity.Product, error)
	GetProductList(product entity.Product) ([]entity.Product, error)
	RemoveProduct(product entity.Product) (*mongo.DeleteResult, error)
	UpdateProduct()
}

const PCollection string = "product"

var (
	repoMongo repository.IMongoRepository = repository.ObjIMongoRepository()
)

func InstProductService() IProductService {
	return &productService{}
}

func (*productService) AddProduct(product entity.Product) (*mongo.InsertOneResult, error) {
	result, err := repoMongo.Create(PCollection, product)

	return result, err
}
func (*productService) AddProductList(products []interface{}) (*mongo.InsertManyResult, error) {
	result, err := repoMongo.CreateMany(PCollection, products)

	return result, err
}
func (*productService) GetProduct(product entity.Product) (entity.Product, error) {
	result := repoMongo.FindOne(PCollection, product)
	err := result.Decode(&product)

	return product, err

}
func (*productService) GetProductList(product entity.Product) ([]entity.Product, error) {
	var productList []entity.Product
	result, err := repoMongo.FindAll(PCollection, product)
	if err != nil {
		log.Println(err)
	}
	err = result.All(context.TODO(), &productList)

	return productList, err
}

func (*productService) RemoveProduct(product entity.Product) (*mongo.DeleteResult, error) {
	result, err := repoMongo.DeleteDocument(PCollection, product)
	return result, err
}
func (*productService) UpdateProduct() {}

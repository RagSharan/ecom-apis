package service

import (
	"context"
	"log"

	"github.com/ragsharan/ecom-apis/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const BUNDLE string = "bundle"

type bundleService struct{}

type IBundleService interface {
	AddBundle(bundle entity.Bundle) (*mongo.InsertOneResult, error)
	GetBundle(bundle entity.Bundle) (entity.Bundle, error)
	GetBundleList(bundle entity.Bundle) ([]entity.Bundle, error)

	RemoveBundle(bundle entity.Bundle) (*mongo.DeleteResult, error)
	UpdateBundle()
}

func InstBundleService() IBundleService {
	return &bundleService{}
}

func (*bundleService) AddBundle(bundle entity.Bundle) (*mongo.InsertOneResult, error) {
	result, err := repoMongo.Create(BUNDLE, bundle)
	return result, err

}
func (*bundleService) GetBundle(bundle entity.Bundle) (entity.Bundle, error) {
	result := repoMongo.FindOne(BUNDLE, bundle)
	err := result.Decode(&bundle)
	return bundle, err
}
func (*bundleService) GetBundleList(bundle entity.Bundle) ([]entity.Bundle, error) {
	var bundleList []entity.Bundle
	result, err := repoMongo.FindList(BUNDLE, bundle)
	if err != nil {
		log.Println(err)
	}
	err = result.All(context.TODO(), &bundleList)
	return bundleList, err
}

func (*bundleService) RemoveBundle(bundle entity.Bundle) (*mongo.DeleteResult, error) {
	result, err := repoMongo.DeleteDocument(BUNDLE, bundle)
	return result, err
}
func (*bundleService) UpdateBundle() {}

package service

import (
	"context"
	"log"

	"github.com/ragsharan/ecom-apis/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const REVIEW string = "review"

type reviewService struct{}

type IReviewService interface {
	AddReview(review entity.Review) (*mongo.InsertOneResult, error)
	GetReview(review entity.Review) (entity.Review, error)
	GetReviewList(review entity.Review) ([]entity.Review, error)
	RemoveReview(review entity.Review) (*mongo.DeleteResult, error)
	UpdateReview()
}

func InstReviewService() IReviewService {
	return &reviewService{}
}

func (*reviewService) AddReview(review entity.Review) (*mongo.InsertOneResult, error) {
	result, err := repoMongo.Create(REVIEW, review)
	return result, err
}

func (*reviewService) GetReview(review entity.Review) (entity.Review, error) {
	result := repoMongo.FindOne(REVIEW, review)
	err := result.Decode(&review)
	return review, err
}
func (*reviewService) GetReviewList(review entity.Review) ([]entity.Review, error) {
	var reviewList []entity.Review
	result, err := repoMongo.FindAll(REVIEW, review)
	if err != nil {
		log.Println(err)
	}
	err = result.All(context.TODO(), &reviewList)
	return reviewList, err
}

func (*reviewService) RemoveReview(review entity.Review) (*mongo.DeleteResult, error) {
	result, err := repoMongo.DeleteDocument(REVIEW, review)
	return result, err
}

func (*reviewService) UpdateReview() {}

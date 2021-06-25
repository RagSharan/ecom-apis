package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ragsharan/ecom-apis/entity"
	"github.com/ragsharan/ecom-apis/service"
)

type reviewControl struct{}
type IReviewControl interface {
	AddReview(res http.ResponseWriter, req *http.Request)
	GetReview(res http.ResponseWriter, req *http.Request)
	GetReviewList(res http.ResponseWriter, req *http.Request)
	RemoveReview(res http.ResponseWriter, req *http.Request)
	UpdateReview(res http.ResponseWriter, req *http.Request)
}

var (
	reviewService service.IReviewService = service.InstReviewService()
)

func InstReviewControl() IReviewControl {
	return &reviewControl{}
}

func (*reviewControl) AddReview(res http.ResponseWriter, req *http.Request) {
	var review entity.Review
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &review)
	result, err := reviewService.AddReview(review)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*reviewControl) GetReview(res http.ResponseWriter, req *http.Request) {
	var review entity.Review
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &review)
	result, err := reviewService.GetReview(review)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*reviewControl) GetReviewList(res http.ResponseWriter, req *http.Request) {
	var review entity.Review
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &review)
	result, err := reviewService.GetReviewList(review)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*reviewControl) RemoveReview(res http.ResponseWriter, req *http.Request) {
	var review entity.Review
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &review)
	result, err := reviewService.RemoveReview(review)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*reviewControl) UpdateReview(res http.ResponseWriter, req *http.Request) {}

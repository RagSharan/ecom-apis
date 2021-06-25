package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ragsharan/ecom-apis/entity"
	"github.com/ragsharan/ecom-apis/service"
)

type bundleControl struct{}
type IBundleControl interface {
	CreateBundle(res http.ResponseWriter, req *http.Request)
	GetBundle(res http.ResponseWriter, req *http.Request)
	GetBundleList(res http.ResponseWriter, req *http.Request)

	RemoveBundle(res http.ResponseWriter, req *http.Request)
	RemoveFromBundle(res http.ResponseWriter, req *http.Request)
	AddInBundle(res http.ResponseWriter, req *http.Request)
	UpdateBundle(res http.ResponseWriter, req *http.Request)
}

var (
	bundleService service.IBundleService = service.InstBundleService()
)

func InstBundleControl() IBundleControl {
	return &bundleControl{}
}

func (*bundleControl) CreateBundle(res http.ResponseWriter, req *http.Request) {
	var bundle entity.Bundle
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &bundle)
	result, err := bundleService.AddBundle(bundle)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*bundleControl) GetBundle(res http.ResponseWriter, req *http.Request) {
	var bundle entity.Bundle
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &bundle)
	result, err := bundleService.GetBundle(bundle)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*bundleControl) GetBundleList(res http.ResponseWriter, req *http.Request) {
	var bundle entity.Bundle
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &bundle)
	result, err := bundleService.GetBundleList(bundle)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}

func (*bundleControl) RemoveBundle(res http.ResponseWriter, req *http.Request) {
	var bundle entity.Bundle
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &bundle)
	result, err := bundleService.RemoveBundle(bundle)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}
	res.WriteHeader(http.StatusAccepted)
	json.NewEncoder(res).Encode(result)
}
func (*bundleControl) RemoveFromBundle(res http.ResponseWriter, req *http.Request) {}
func (*bundleControl) AddInBundle(res http.ResponseWriter, req *http.Request)      {}
func (*bundleControl) UpdateBundle(res http.ResponseWriter, req *http.Request)     {}

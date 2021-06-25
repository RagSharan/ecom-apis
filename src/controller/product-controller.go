package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ragsharan/ecom-apis/entity"
	"github.com/ragsharan/ecom-apis/service"
)

type productControl struct{}

type IProductControl interface {
	AddProduct(res http.ResponseWriter, req *http.Request)
	AddProducts(res http.ResponseWriter, req *http.Request)
	GetProduct(res http.ResponseWriter, req *http.Request)
	GetProducts(res http.ResponseWriter, req *http.Request)
	RemoveProduct(res http.ResponseWriter, req *http.Request)
	UpdateProduct(res http.ResponseWriter, req *http.Request)
	UpdateProductList(res http.ResponseWriter, req *http.Request)
}

var (
	prodService service.IProductService = service.InstProductService()
)

func InstProductControl() IProductControl {
	return &productControl{}
}

func (*productControl) AddProduct(res http.ResponseWriter, req *http.Request) {
	var product entity.Product
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &product)

	result, err := prodService.AddProduct(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(result)
}

func (*productControl) AddProducts(res http.ResponseWriter, req *http.Request) {
	var productList []interface{}
	var tempList []entity.Product
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &tempList)
	for _, val := range tempList {
		productList = append(productList, val)
	}
	result, err := prodService.AddProductList(productList)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(result)
}

func (*productControl) GetProduct(res http.ResponseWriter, req *http.Request) {
	var product entity.Product
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &product)
	product, err = prodService.GetProduct(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(product)

}
func (*productControl) GetProducts(res http.ResponseWriter, req *http.Request) {
	var product entity.Product
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &product)
	result, err := prodService.GetProductList(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(result)
}

func (*productControl) RemoveProduct(res http.ResponseWriter, req *http.Request) {
	var product entity.Product
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &product)
	result, err := prodService.RemoveProduct(product)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(result)
}

func (*productControl) UpdateProduct(res http.ResponseWriter, req *http.Request)     {}
func (*productControl) UpdateProductList(res http.ResponseWriter, req *http.Request) {}

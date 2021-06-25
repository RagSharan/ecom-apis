package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ragsharan/ecom-apis/entity"
	"github.com/ragsharan/ecom-apis/service"
)

type orderControl struct{}
type IOrderControl interface {
	GetOrder(res http.ResponseWriter, req *http.Request)
	GetUserOrders(res http.ResponseWriter, req *http.Request)
	AddOrder(res http.ResponseWriter, req *http.Request)
	UpdateOrder(res http.ResponseWriter, req *http.Request)
}

var (
	orderService service.IOrderService = service.InstOrderService()
)

func InstOrderControl() IOrderControl {
	return &orderControl{}
}

func (*orderControl) GetOrder(res http.ResponseWriter, req *http.Request) {
	var order entity.Order
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &order)
	order, err = orderService.GetOrder(order)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(order)
}
func (*orderControl) GetUserOrders(res http.ResponseWriter, req *http.Request) {
	var order entity.Order
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &order)
	result, err := orderService.GetOrderList(order)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(result)
}
func (*orderControl) AddOrder(res http.ResponseWriter, req *http.Request) {
	var order entity.Order
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(data, &order)
	result, err := orderService.CreateOrder(order)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
	}
	json.NewEncoder(res).Encode(result)
}
func (*orderControl) UpdateOrder(res http.ResponseWriter, req *http.Request) {}

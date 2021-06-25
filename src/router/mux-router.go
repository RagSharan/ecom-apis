package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() IRouter {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) PUT(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) DELETE(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}
func (*muxRouter) SERVE(port string) {
	fmt.Printf("MUX http server running on port  %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

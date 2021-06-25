package router

import "github.com/ragsharan/ecom-apis/controller"

type mapping struct{}

type IMapping interface {
	UriMappings()
}

var (
	httpRouter     IRouter                    = NewMuxRouter()
	productControl controller.IProductControl = controller.InstProductControl()
)

func InstMapping() IMapping {
	return mapping{}
}
func (mapping) UriMappings() {
	productMapping()
	orderMapping()

}
func productMapping() {
	httpRouter.GET("/product", productControl.GetProduct)
	httpRouter.GET("/postList", productControl.GetProducts)
	httpRouter.POST("/product", productControl.AddProduct)
	httpRouter.POST("/postList", productControl.AddProducts)
	httpRouter.PUT("/product", productControl.UpdateProduct)
	httpRouter.PUT("/postList", productControl.UpdateProductList)
	httpRouter.DELETE("/product", productControl.RemoveProduct)
}

func orderMapping() {

}

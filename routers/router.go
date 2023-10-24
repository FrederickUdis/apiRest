package routers

import (
	"apiRest/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/productos", &controllers.ProductoController{}, "post:Post;get:Get")
	beego.Router("/productos/:id:int", &controllers.ProductoController{}, "put:Put;delete:Delete")
	beego.Router("/consumidores", &controllers.ConsumidorController{}, "post:Post;get:Get")
	beego.Router("/consumidores/:id:int", &controllers.ConsumidorController{}, "put:Put;delete:Delete")
	beego.Router("/ordenes", &controllers.OrdenController{}, "post:Post;get:Get")
	beego.Router("/ordenes/:id:int", &controllers.OrdenController{}, "put:Put;delete:Delete")
}

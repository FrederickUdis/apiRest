package controllers

import (
	"apiRest/models"
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type OrdenController struct {
	web.Controller
}

func (c *OrdenController) Get() {
	o := orm.NewOrm()
	var ordenes []models.Orden

	// Consulta para obtener todas las órdenes
	qs := o.QueryTable("orden")

	// Eager loading para cargar las relaciones de Consumidor y Producto
	qs = qs.RelatedSel("Consumidor").RelatedSel("Producto")

	// Ejecutar la consulta
	qs.All(&ordenes)

	// Ahora, 'ordenes' debería tener las relaciones de Consumidor y Producto completamente cargadas
	c.Data["json"] = ordenes
	c.ServeJSON()
	log.Println("Se obtuvieron todas las órdenes")
}

// Post implementa la lógica para crear una nueva orden
func (c *OrdenController) Post() {
	log.Printf("Bandera entro en el post")
	var orden models.Orden
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orden)
	if err != nil {
		log.Printf("Error al deserializar el body: %s\n", err.Error())
		c.Data["json"] = "Error al procesar la solicitud"
		c.ServeJSON()
		return
	}
	log.Printf("Body recibido para POST: ID: %d, Consumidor: %+v, Producto: %+v, Cantidad: %d, Total: %f\n", orden.ID, *orden.Consumidor, *orden.Producto, orden.Cantidad, orden.Total)
	o := orm.NewOrm()
	_, err = o.Insert(&orden)
	if err != nil {
		log.Printf("Error al guardar la orden: %s\n", err.Error())
		c.Data["json"] = "Error al guardar la orden"
		c.ServeJSON()
		return
	}

	c.Data["json"] = orden
	c.ServeJSON()
	log.Printf("Se creó una nueva orden con ID: %d\n", orden.ID)
}

// Put implementa la lógica para actualizar una orden existente
func (c *OrdenController) Put() {
	id, _ := c.GetInt64(":id")
	o := orm.NewOrm()
	orden := models.Orden{ID: id}
	if o.Read(&orden) == nil {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &orden)
		if err != nil {
			log.Printf("Error al deserializar el body: %s\n", err.Error())
			c.Data["json"] = "Error al procesar la solicitud"
			c.ServeJSON()
			return
		}
		log.Printf("Body recibido para PUT (ID %d): ID: %d, Consumidor: %+v, Producto: %+v, Cantidad: %d, Total: %f\n", id, orden.ID, *orden.Consumidor, *orden.Producto, orden.Cantidad, orden.Total)
		_, err = o.Update(&orden)
		if err != nil {
			log.Printf("Error al actualizar la orden con ID %d: %s\n", id, err.Error())
			c.Data["json"] = "Error al actualizar la orden"
			c.ServeJSON()
			return
		}

		c.Data["json"] = orden
		log.Printf("Se actualizó la orden con ID: %d\n", orden.ID)
	} else {
		c.Data["json"] = "Orden no encontrada"
		log.Printf("Error al actualizar: Orden con ID %d no encontrada\n", id)
	}
	c.ServeJSON()
}

// Delete implementa la lógica para eliminar una orden
func (c *OrdenController) Delete() {
	id, _ := c.GetInt64(":id")
	o := orm.NewOrm()
	orden := models.Orden{ID: id}
	if num, err := o.Delete(&orden); err == nil {
		c.Data["json"] = map[string]int64{"num": num}
		log.Printf("Se eliminó la orden con ID: %d\n", id)
	} else {
		c.Data["json"] = err.Error()
		log.Printf("Error al eliminar la orden con ID: %d. Error: %s\n", id, err.Error())
	}
	c.ServeJSON()
}

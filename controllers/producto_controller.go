package controllers

import (
	"apiRest/models"
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type ProductoController struct {
	web.Controller
}

// Get implementa la lógica para obtener productos
func (c *ProductoController) Get() {
	o := orm.NewOrm()
	var productos []models.Producto
	o.QueryTable("producto").All(&productos)
	c.Data["json"] = productos
	c.ServeJSON()
}

// Post implementa la lógica para crear un nuevo producto
func (c *ProductoController) Post() {
	var producto models.Producto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &producto); err != nil {
		log.Printf("Error al deserializar el cuerpo de la solicitud: %v", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Error al procesar los datos enviados"
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	_, err := o.Insert(&producto)
	if err != nil {
		log.Printf("Error al insertar el producto en la base de datos: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = "Error al guardar el producto"
		c.ServeJSON()
		return
	}
	c.Data["json"] = producto
	c.ServeJSON()
}

// Put implementa la lógica para actualizar un producto existente
func (c *ProductoController) Put() {
	id, _ := c.GetInt64(":id")
	log.Printf("Actualizando producto con ID: %d\n", id)
	o := orm.NewOrm()
	producto := models.Producto{ID: id}
	if err := o.Read(&producto); err == nil {
		json.Unmarshal(c.Ctx.Input.RequestBody, &producto)
		if _, err := o.Update(&producto); err != nil {
			log.Printf("Error al actualizar producto con ID: %d. Error: %v\n", id, err)
		} else {
			log.Printf("Producto con ID: %d actualizado exitosamente\n", id)
			c.Data["json"] = producto
		}
	} else {
		log.Printf("Producto con ID: %d no encontrado. Error: %v\n", id, err)
		c.Data["json"] = "Producto no encontrado"
	}
	c.ServeJSON()
}

// Delete implementa la lógica para eliminar un producto
func (c *ProductoController) Delete() {
	id, _ := c.GetInt64(":id")
	o := orm.NewOrm()
	producto := models.Producto{ID: id}
	if num, err := o.Delete(&producto); err == nil {
		c.Data["json"] = map[string]int64{"num": num}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

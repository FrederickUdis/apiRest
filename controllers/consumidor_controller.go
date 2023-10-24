package controllers

import (
	"apiRest/models"
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type ConsumidorController struct {
	web.Controller
}

// Get implementa la lógica para obtener consumidores
func (c *ConsumidorController) Get() {
	o := orm.NewOrm()
	var consumidores []models.Consumidor
	o.QueryTable("consumidor").All(&consumidores)
	c.Data["json"] = consumidores
	c.ServeJSON()
}

// Post implementa la lógica para crear un nuevo consumidor
func (c *ConsumidorController) Post() {
	var consumidor models.Consumidor
	json.Unmarshal(c.Ctx.Input.RequestBody, &consumidor)
	o := orm.NewOrm()
	o.Insert(&consumidor)
	c.Data["json"] = consumidor
	c.ServeJSON()
}

// Put implementa la lógica para actualizar un consumidor existente
func (c *ConsumidorController) Put() {
	id, _ := c.GetInt64(":id")
	o := orm.NewOrm()
	consumidor := models.Consumidor{ID: id}
	if o.Read(&consumidor) == nil {
		json.Unmarshal(c.Ctx.Input.RequestBody, &consumidor)
		o.Update(&consumidor)
		c.Data["json"] = consumidor
	} else {
		c.Data["json"] = "Consumidor no encontrado"
	}
	c.ServeJSON()
}

// Delete implementa la lógica para eliminar un consumidor
func (c *ConsumidorController) Delete() {
	id, _ := c.GetInt64(":id")
	o := orm.NewOrm()
	consumidor := models.Consumidor{ID: id}
	if num, err := o.Delete(&consumidor); err == nil {
		c.Data["json"] = map[string]int64{"num": num}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

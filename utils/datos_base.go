package utils

import (
	"apiRest/models"

	"github.com/beego/beego/v2/client/orm"
)

func InsertarDatosPrueba() {
	o := orm.NewOrm()

	// Datos de prueba para Consumidor
	consumidor1 := models.Consumidor{
		Nombre: "Juan Pérez",
		Email:  "juan.perez@example.com",
	}
	o.Insert(&consumidor1)

	consumidor2 := models.Consumidor{
		Nombre: "María García",
		Email:  "maria.garcia@example.com",
	}
	o.Insert(&consumidor2)

	// Datos de prueba para Producto
	producto1 := models.Producto{
		Descripcion: "Laptop",
		Precio:      1000.50,
		Stock:       10,
	}
	o.Insert(&producto1)

	producto2 := models.Producto{
		Descripcion: "Teléfono",
		Precio:      500.25,
		Stock:       20,
	}
	o.Insert(&producto2)

	// Datos de prueba para Orden
	orden1 := models.Orden{
		Consumidor: &consumidor1,
		Producto:   &producto1,
		Cantidad:   2,
		Total:      2001.00,
	}
	o.Insert(&orden1)

	orden2 := models.Orden{
		Consumidor: &consumidor2,
		Producto:   &producto2,
		Cantidad:   1,
		Total:      500.25,
	}
	o.Insert(&orden2)
}

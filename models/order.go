package models

type Orden struct {
	ID         int64       `orm:"auto;pk"`
	Consumidor *Consumidor `orm:"rel(fk)"` // Relación con Consumidor
	Producto   *Producto   `orm:"rel(fk)"` // Relación con Producto
	Cantidad   int
	Total      float64
}

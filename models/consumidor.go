package models

type Consumidor struct {
	ID      int64    `orm:"auto;pk"`
	Nombre  string   `orm:"size(100)"`
	Email   string   `orm:"size(100);unique"`
	Ordenes []*Orden `orm:"reverse(many)"` // Relación inversa
}

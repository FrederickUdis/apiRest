package models

type Producto struct {
	ID          int64    `json:"ID"`
	Descripcion string   `json:"Descripcion"`
	Precio      float64  `json:"Precio"`
	Stock       int      `json:"Stock"`
	Ordenes     []*Orden `orm:"reverse(many)" json:"Ordenes"`
}

package models


type Concepto struct {
	Id             int    `orm:"column(id);pk"`
	Codigo         string `orm:"column(codigo)"`
	Nombre         string `orm:"column(nombre)"`
	TipoAfectacion int16  `orm:"column(tipo_afectacion)"`
	Rubro          *Rubro `orm:"column(rubro);rel(fk)"`
}

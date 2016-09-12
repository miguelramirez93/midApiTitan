package models


type CategoriaNovedad struct {
	Id          int    `orm:"column(id);pk"`
	Nombre      string `orm:"column(nombre)"`
	Descripcion string `orm:"column(descripcion)"`
	Estado      string `orm:"column(estado)"`
}

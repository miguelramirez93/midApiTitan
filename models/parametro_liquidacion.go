package models



type ParametroLiquidacion struct {
	Id          int                 `orm:"column(id);pk"`
	Nombre      string              `orm:"column(nombre)"`
	Descripcion string              `orm:"column(descripcion)"`
	Simbolo     string              `orm:"column(simbolo)"`
	Valor       float64             `orm:"column(valor)"`
	Estado      string              `orm:"column(estado)"`
	IdCategoria *CategoriaParametro `orm:"column(id_categoria);rel(fk)"`
}

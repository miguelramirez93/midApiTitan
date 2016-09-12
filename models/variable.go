package models



type Variable struct {
	Id      int    `orm:"column(id);pk"`
	Nombre  string `orm:"column(nombre)"`
	Simbolo string `orm:"column(simbolo)"`
	Valor   string `orm:"column(valor)"`
}

package models



type Nomina struct {
	Id          int    `pk;orm:"column(id)"`
	Vinculacion int    `orm:"column(vinculacion);null"`
	Nombre      string `orm:"column(nombre)"`
	Descripcion string `orm:"column(descripcion)"`
	TipoNomina  string `orm:"column(tipo_nomina)"`
	Estado      string `orm:"column(estado)"`
	Periodo     string `orm:"column(periodo);null"`
}

package models



type DetalleNovedad struct {
	Id       int     `pk;orm:"column(novedad)"`
	Persona  int64   `orm:"column(persona)"`
	Estado   string  `orm:"column(estado)"`
	Vigencia float64 `orm:"column(vigencia);null"`
	Valor    float64 `orm:"column(valor);null"`
	Cuenta   string  `orm:"column(cuenta);null"`
}

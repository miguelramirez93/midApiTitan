package models



type DetalleNovedad struct {
	Id       int     `pk;orm:"column(id);serial"`
	Persona  int64   `orm:"column(persona)"`
	Estado   string  `orm:"column(estado)"`
	Vigencia int `orm:"column(vigencia);null"`
	Valor    float64 `orm:"column(valor);null"`
	Cuenta   string  `orm:"column(cuenta);null"`
	Novedad  *Novedad `orm:"rel(fk);column(novedad)"`
}

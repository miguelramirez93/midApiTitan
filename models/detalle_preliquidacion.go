package models


type DetallePreliquidacion struct {
	Id       int       `pk;orm:"column(preliquidacion)"`
	Persona  int64     `orm:"column(persona)"`
	Valor    float32   `orm:"column(valor)"`
	Concepto *Concepto `orm:"column(concepto);rel(fk)"`
}

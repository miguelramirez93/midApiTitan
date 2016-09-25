package models

type DetallePreliquidacion struct {
	Preliquidacion *Preliquidacion `orm:"column(preliquidacion);rel(fk)"`
	Persona        int64           `orm:"column(persona)"`
	Valor          string         `orm:"column(valor)"`
	Concepto       *Concepto       `orm:"column(concepto);rel(fk)"`
	Id             int             `orm:"column(id);pk"`
}

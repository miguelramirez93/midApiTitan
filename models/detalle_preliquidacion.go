package models

type DetallePreliquidacion struct {
	Preliquidacion *Preliquidacion `orm:"column(preliquidacion);rel(fk)"`
	Persona        *InformacionProveedor           `orm:"column(persona);rel(fk)"`
	ValorBruto          string         `orm:"column(valor_bruto)"`
	Valor          string         `orm:"column(valor)"`
	Concepto       *Concepto       `orm:"column(concepto);rel(fk)"`
	Id             int             `orm:"column(id);pk"`
}

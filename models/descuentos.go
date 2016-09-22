package models



type Descuentos struct {
	Id                    int                    `pk;orm:"column(id)"`
	Nombre                string                 `orm:"column(nombre);null"`
	Base                  float64                `orm:"column(base);null"`
	Valor                 float64                `orm:"column(valor);null"`
	DetallePreliquidacion *DetallePreliquidacion `orm:"column(detalle_preliquidacion);rel(fk)"`
}

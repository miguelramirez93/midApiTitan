package models


type NovedadAplicada struct {
	Id                    int                    `orm:"column(id);pk"`
	DetallePreliquidacion *DetallePreliquidacion `orm:"column(detalle_preliquidacion);rel(fk)"`
	DetalleNovedad        *DetalleNovedad        `orm:"column(detalle_novedad);rel(fk)"`
}

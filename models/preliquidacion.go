package models

import (

	"time"


)

type Preliquidacion struct {
	Id          int       `pk;orm:"column(id)"`
	Nombre      string    `orm:"column(nombre)"`
	Nomina      *Nomina   `orm:"column(nomina);rel(fk)"`
	IdUsuario   int64     `orm:"column(id_usuario)"`
	Estado      string    `orm:"column(estado)"`
	Fecha       time.Time `orm:"column(fecha);type(date)"`
	Descripcion string    `orm:"column(descripcion);null"`
	FechaInicio time.Time `orm:"column(fecha_inicio);type(date)"`
	FechaFin    time.Time `orm:"column(fecha_fin);type(date)"`
}

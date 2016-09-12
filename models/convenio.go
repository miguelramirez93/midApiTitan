package models

import (

	"time"


)

type Convenio struct {
	Id            int       `orm:"column(NUMERO_PRO);pk"`
	CLAVE         string    `orm:"column(CLAVE);null"`
	NPE           string    `orm:"column(NPE);null"`
	ANIOPRO       int       `orm:"column(ANIO_PRO);null"`
	NOMBRE        string    `orm:"column(NOMBRE);null"`
	OBJETO        string    `orm:"column(OBJETO);null"`
	ENTIDAD       string    `orm:"column(ENTIDAD);null"`
	CODIGOTESORAL int64     `orm:"column(CODIGO_TESORAL);null"`
	FECHAINICIO   time.Time `orm:"column(FECHA_INICIO);type(date);null"`
	FECHAFINAL    time.Time `orm:"column(FECHA_FINAL);type(date);null"`
	SITUACION     string    `orm:"column(SITUACION);null"`
	UNIDAD        string    `orm:"column(UNIDAD);null"`
	ESTADO        string    `orm:"column(ESTADO);null"`
	MODALIDAD     string    `orm:"column(MODALIDAD);null"`
}

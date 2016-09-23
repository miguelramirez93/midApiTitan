package models

import (

	"time"


)

type ContratoGeneral struct {
	Id              string           `orm:"column(numero_contrato);pk"`
	Vigencia                     int              `orm:"column(vigencia)"`
	IdOrdenContrato              int              `orm:"column(id_orden_contrato);null"`
	TipoContrato                 *Parametros      `orm:"column(tipo_contrato);rel(fk);null"`
	UnidadEjecutora              string           `orm:"column(unidad_ejecutora);null"`
	ObjetoContrato               string           `orm:"column(objeto_contrato);null"`
	FechaInicio                  time.Time        `orm:"column(fecha_inicio);type(date);null"`
	FechaFinal                   time.Time        `orm:"column(fecha_final);type(date);null"`
	PlazoEjecucion               int              `orm:"column(plazo_ejecucion);null"`
	FormaPago                    *Parametros      `orm:"column(forma_pago);rel(fk);null"`
	Supervisor                   string           `orm:"column(supervisor);null"`
	ClausulaRegistroPresupuestal bool             `orm:"column(clausula_registro_presupuestal);null"`
	SedeSupervisor               string           `orm:"column(sede_supervisor);null"`
	DependenciaSupervisor        string           `orm:"column(dependencia_supervisor);null"`
	CargoSupervisor              string           `orm:"column(cargo_supervisor);null"`
	SedeSolicitante              string           `orm:"column(sede_solicitante);null"`
	DependenciaSolicitante       string           `orm:"column(dependencia_solicitante);null"`
	NumeroSolicitudNecesidad     int              `orm:"column(numero_solicitud_necesidad);null"`
	NumeroCdp                    int              `orm:"column(numero_cdp);null"`
	EstadoAprobacion             bool             `orm:"column(estado_aprobacion);null"`
	Contratista                  *InformacionProveedor             `orm:"column(contratista);rel(fk)"`
	NombreContratista            string           `orm:"column(nombre_contratista);null"`
	UnidadEjecucion              *Parametros      `orm:"column(unidad_ejecucion);rel(fk);null"`
	ValorContrato                string             `orm:"column(valor_contrato);null"`
	Estado                       int              `orm:"column(estado);null"`
	Justificacion                string           `orm:"column(justificacion);null"`
	DescripcionFormaPago         string           `orm:"column(descripcion_forma_pago);null"`
	Condiciones                  string           `orm:"column(condiciones);null"`
}

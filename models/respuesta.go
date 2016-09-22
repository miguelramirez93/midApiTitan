package models

type Respuesta struct {
	Nombre_Cont         string
	Valor_bruto string
	Novedades  string
	Retenciones  string
	Valor_neto  string
	Descuentos *[]Descuentos
}
type FormatoPreliqu struct {
	Contrato   *ContratoGeneral
	Respuesta *Respuesta
}

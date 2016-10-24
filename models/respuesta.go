package models

type Respuesta struct {
	Nombre_Cont         string
	NumDocumento         int64
	Valor_bruto string
	Valor_neto  string
	Descuentos *[]Descuentos
	Novedades *[]DetalleNovedad
}
type FormatoPreliqu struct {
	//Contrato   *ContratoGeneral
	Respuesta *Respuesta
}

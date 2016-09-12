package models

type Respuesta struct {
	Nombre_Cont         string
	Valor_contrato      string
	Valor_pago string
  Valor_descuento_salud string
  Valor_descuento_pension string
	Valores_finales  string
}
type FormatoPreliqu struct {
	Contrato   *ContratoGeneral
	Respuesta *Respuesta
}

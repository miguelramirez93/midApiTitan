package golog

import (
  "fmt"
  models "github.com/miguelramirez93/midApiTitan/models"
  . "github.com/mndrix/golog"
)

func CargarReglas(reglas string) (rest []models.Respuesta) {

  m := NewMachine().Consult(reglas)
  var control int = 0
  contratos := m.ProveAll(`valor_contrato(X,Y).`)
  var miArreglo = make([]models.Respuesta, len(contratos)+1)
  miArreglo[0].Nombre_Cont = "Nombre contratista"
  miArreglo[0].Valor_contrato = "Valor contrato"
  miArreglo[0].Valor_pago = "Valor pago"
  miArreglo[0].Valor_descuento_salud = "Valor descuento salud"
  miArreglo[0].Valor_descuento_pension = "Valor descuento pension"
  miArreglo[0].Valores_finales = "Valor final"
  control++
  for _, solution := range contratos {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_contrato = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }

  control = 1
  pagos_brutos := m.ProveAll(`valor_pago(X,Y).`)
  for _, solution := range pagos_brutos {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_pago = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }

  control = 1
  pagos_salud := m.ProveAll(`valor_descuento_salud(X,Y).`)
  for _, solution := range pagos_salud {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_descuento_salud = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }

  control = 1
  pagos_pension := m.ProveAll(`valores_finales(X,Y).`)
  for _, solution := range pagos_pension {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valores_finales = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }

  return miArreglo

}

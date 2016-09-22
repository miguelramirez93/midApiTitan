package golog

import (
  "fmt"
  "strconv"
  models "github.com/miguelramirez93/midApiTitan/models"
  . "github.com/mndrix/golog"
)

func CargarReglas(reglas string, periodo string) (rest []models.Respuesta) {
//******QUITAR ARREGLO, DEJAR UNA SOLA VARIABLE PARA LAS REGLAS ******
  m := NewMachine().Consult(reglas)
  var control int = 0



  /*contratos := m.ProveAll(`valor_contrato(X,Y).`)
  for _, solution := range contratos {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_contrato = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }*/

  /*control = 1
  pagos_brutos := m.ProveAll(`valor_pago(X,Y).`)
  for _, solution := range pagos_brutos {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_pago = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }*/

  /*control = 1
  pagos_salud := m.ProveAll(`valor_descuento_salud(X,Y).`)
  for _, solution := range pagos_salud {
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_descuento_salud = fmt.Sprintf("%s", solution.ByName_("Y"))
    control++
  }*/


  pagos_pension := m.ProveAll("valor_pago_neto(X,Y,"+periodo+",V,L,L2).")
  var miArreglo = make([]models.Respuesta, len(pagos_pension))
  for _, solution := range pagos_pension {
      miArreglo[control].Valor_neto = fmt.Sprintf("%s", solution.ByName_("Y"))
    miArreglo[control].Nombre_Cont = fmt.Sprintf("%s", solution.ByName_("X"))
    miArreglo[control].Valor_bruto  = fmt.Sprintf("%s", solution.ByName_("V"))
    miArreglo[control].Novedades = fmt.Sprintf("%s", solution.ByName_("L"))
    miArreglo[control].Retenciones = fmt.Sprintf("%s", solution.ByName_("L2"))


    control++
  }
  control = 0
  descuentos := m.ProveAll("retencion(X,Y,"+periodo+",B,N).")
  var lista_descuentos = make([]models.Descuentos, len(descuentos))
  for _, solution := range descuentos {
    lista_descuentos[control].Nombre = fmt.Sprintf("%s", solution.ByName_("N"))
    lista_descuentos[control].Base,_ = strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("B")), 64)
    lista_descuentos[control].Valor,_ = strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("Y")),64)
    control++
  }
    miArreglo[0].Descuentos = &lista_descuentos
  return miArreglo

}

package main

import (
     "fmt"
     . "github.com/mndrix/golog"
     //"strings"
     "strconv"
)

func main(){
m := NewMachine().Consult(`
  valor_pago(prueba,2016,3143915).
  factor(prueba2, descuento, porcentaje, salud, 0.04, 2016).
  factor(prueba2, descuento, porcentaje, pension, 0.04, 2016).
  valor_uvt(29753,2016).
  tarifa(ret383,0,95,0,0).
  tarifa(ret383,95,150,0.19,0).
  tarifa(ret383,150,360,0.28,10).
  tarifa(ret383,360,999999999999,0.33,69).
  tarifa(ret384,0,128.96,0,0).
  tarifa(ret384,128.96,132.36,0.09,0).
  tarifa(ret384,132.36,135.75,0.09,0).
  tarifa(ret384,135.75,139.14,0.09,0).
  tarifa(ret384,139.14,142.54,0.09,0).
  tarifa(ret384,142.54,145.93,0.1,0).
  tarifa(ret384,145.93,149.32,0.2,0).
  tarifa(ret384,149.32,152.72,0.2,0).
  tarifa(ret384,152.72,156.11,0.21,0).
  tarifa(ret384,156.11,159.51,0.4,0).
  tarifa(ret384,159.51,162.9,0.41,0).
  tarifa(ret384,162.9,166.29,0.1,0).
  tarifa(ret384,162.9,166.29,0.41,0).
  tarifa(ret384,166.29,169.69,0.7,0).
  tarifa(ret384,169.69,176.47,0.73,0).
  tarifa(ret384,176.47,183.26,1.15,0).
  tarifa(ret384,183.26,190.05,1.19,0).
  tarifa(ret384,190.05,196.84,1.65,0).
  tarifa(ret384,196.84,203.62,2.14,0).
  tarifa(ret384,203.62,210.41,2.21).
  tarifa(ret384,210.41,217.2,2.96,0).
  tarifa(ret384,217.2,230.77,3.87,0).
  tarifa(ret384,230.77,237.56,4.63,0).
  tarifa(ret384,237.56,244.35,5.06,0).
  tarifa(ret384,244.35,257.92,5.96,0).
  tarifa(ret384,257.92,264.71,6.44,0).
  tarifa(ret384,264.71,271.5,6.93,0).
  tarifa(ret384,271.5,278.29,7.44,0).
  tarifa(ret384,278.29,285.07,7.96,0).
  tarifa(ret384,285.07,291.86,8.5,0).
  tarifa(ret384,291.86,298.65,9.05,0).
  tarifa(ret384,298.65,305.44,9.62,0).
  tarifa(ret384,305.44,312.22,10.21,0).
  tarifa(ret384,312.22,319.01,10.81,0).
  tarifa(ret384,319.01,325.8,11.43,0).
  tarifa(ret384,325.8,332.59,12.07,0).
  tarifa(ret384,332.59,339.37,12.71,0).
  tarifa(ret384,339.37,356.34,14.06,0).
  tarifa(ret384,356.34,373.31,15.83,0).
  tarifa(ret384,373.31,390.28,17.69,0).
  tarifa(ret384,390.28,407.25,19.65,0).
  tarifa(ret384,407.25,424.22,21.69,0).
  tarifa(ret384,424.22,441.19,23.84,0).
  tarifa(ret384,441.19,458.16,26.07,0).
  tarifa(ret384,458.16,475.12,28.39,0).
  tarifa(ret384,475.12,492.09,30.8,0).
  tarifa(ret384,492.09,509.06,33.29,0).
  tarifa(ret384,509.06,526.03,35.87,0).
  tarifa(ret384,526.03,543,38.54,0).
  tarifa(ret384,543,559.97,41.29,0).
  tarifa(ret384,559.97,576.94,44.11,0).
  tarifa(ret384,576.94,593.9,47.02,0).
  tarifa(ret384,593.9,610.87,50,0).
  tarifa(ret384,610.87,627.84,53.06,0).
  tarifa(ret384,627.84,644.81,56.2,0).
  tarifa(ret384,644.81,661.78,59.4,0).
  tarifa(ret384,661.78,678.75,62.68,0).
  tarifa(ret384,678.75,695.72,66.02,0).
  tarifa(ret384,695.72,712.69,69.43,0).
  tarifa(ret384,712.69,729.65,72.9,0).
  tarifa(ret384,729.65,746.62,76.43,0).
  tarifa(ret384,746.62,763.59,80.03,0).
  tarifa(ret384,763.59,780.56,83.68,0).
  tarifa(ret384,780.56,797.53,87.39,0).
  tarifa(ret384,797.53,814.5,91.15,0).
  tarifa(ret384,814.5,831.47,94.96,0).
  tarifa(ret384,831.47,848.44,98.81,0).
  tarifa(ret384,848.44,865.4,102.72,0).
  tarifa(ret384,865.4,882.37,106.67,0).
  tarifa(ret384,882.37,899.34,110.65,0).
  tarifa(ret384,899.34,916.31,114.68,0).
  tarifa(ret384,916.31,933.28,118.74,0).
  tarifa(ret384,933.28,950.25,122.84,0).
  tarifa(ret384,950.25,967.22,126.96,0).
  tarifa(ret384,967.22,984.19,131.11,0).
  tarifa(ret384,984.19,1001.15,135.29,0).
  tarifa(ret384,1001.15,1018.12,139.49,0).
  tarifa(ret384,1018.12,1035.09,143.71,0).
  tarifa(ret384,1035.09,1052.06,147.94,0).
  tarifa(ret384,1052.06,1069.03,152.19,0).
  tarifa(ret384,1069.03,1086,156.45,0).
  tarifa(ret384,1086,1102.97,160.72,0).
  tarifa(ret384,1102.97,1119.93,164.99,0).
  tarifa(ret384,1119.93,1136.92,169.26,0).
  tarifa(reteIca,9.66,1000,0,0).
  tarifa(estampillaUD,0.01,0,0,0).
  tarifa(proCultura,0.005,0,0,0).
  tarifa(adultoMayor,0.005,0,0,0).
  /*valor_hora(X,Y):- valor_contrato(X,V),horas(X,H),Y is V/H.*/
  base_desc_ley(X,Y,P,rete383):- valor_pago(X,P,V),R is (V * 0.25) rnd 0 , S is V - R,Y is (S rnd 0).
  base_desc_ley(X,Y,P,rete384):- valor_pago(X,P,Y).
  base_desc_ley(X,Y,P,secrHacienda):- base_desc_ley(X,Y,P,rete383).
  descuento_ley(X,Y,P,V,rete383):- base_desc_ley(X,V,P,rete383),valor_uvt(U,P),T is V / U ,tarifa(ret383,I,S,C,A),T@>I,T@=<S,R is (((T-I)*C)+A)*U,Y is (R rnd 0).
  descuento_ley(X,Y,P,V,rete384):- base_desc_ley(X,V,P,rete384),valor_uvt(U,P),T is V / U ,tarifa(ret384,I,S,C,A), (T@>I,T@=<S)-> (R is (C * U),Y is (R rnd 0)); (R is (((0.27*T)-135.17) * U) ,Y is (R rnd 0)).
  descuento_ley(X,Y,P,V,reteIca):- base_desc_ley(X,V,P,'secrHacienda'),tarifa(reteIca,I,S,C,A),R is (V*I/S),Y is (R rnd 0).
  descuento_ley(X,Y,P,V,estampillaUD):-base_desc_ley(X,V,P,'secrHacienda'),tarifa(estampillaUD,I,S,C,A),R is (V*I),Y is (R rnd 0).
  descuento_ley(X,Y,P,V,proCultura):-base_desc_ley(X,V,P,'secrHacienda'),tarifa(proCultura,I,S,C,A),R is (V*I),Y is (R rnd 0).
  descuento_ley(X,Y,P,V,adultoMayor):-base_desc_ley(X,V,P,'secrHacienda'),tarifa(adultoMayor,I,S,C,A),R is (V*I),Y is (R rnd 0).
  valores(X,T,P,L):-findall((X, Y, N, Z, R),((factor(X,T,Y,N,Z,P),Y==porcentaje,valor_pago(X,P,V),R is P * Z)),L).
  valores(X,T,P,L):-findall((X, Y, N, Z, R),((factor(X,T,Y,N,Z,P),Y==fijo,R is Z)),L).
  descuentos_ley(X,T,P,L):-findall((X, T, P,Z, R),(descuento_ley(X,R,P,Z,T)),L).
  unir([], Cs, Cs).
  unir([A|As],Bs,[A|Cs]):-unir(As, Bs, Cs).
  total_descuentos([], 0).
  total_descuentos([(X, _, _, _, R)|Xs], S):-total_descuentos(Xs, S2),S is S2 + R.
  valor_pago_neto(X,Y,P,V,L,L2):-valor_pago(X,P,V),valores(X,descuento,P,L),descuentos_ley(X,T,P,L2),unir(L,L2,LS),total_descuentos(LS,D),Y is V - D.
  residuo(N,R):-N @< 1, R is N.
  residuo(N,R):-N @>= 1, N1 is N - 1, residuo(N1,R).
  redondeo(N,R):-residuo(N,X),X @>= 0.5,R is N - X + 1.
  redondeo(N,R):-residuo(N,X),X @< 0.5,R is N - X.
`)
/**
* factor(docente, descuento|suma, porcentaje|fijo, nombre del factor, valor del factor, periodo)
*/
descuentos := m.ProveAll(`valor_pago_neto(prueba,Y,2016,V,L,L2).`)
for _, solution := range descuentos {
    //fmt.Printf("%s -> %s -> %s -> %s\n", solution.ByName_("X"), solution.ByName_("Y"), solution.ByName_("N"), solution.ByName_("Z"))
    Valor,_ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("Y")), 64)
    fmt.Printf("%s\n","Neto")
    fmt.Printf("%.3f\n",Valor)
    fmt.Printf("%s\n","Descuentos")
    fmt.Printf("%s\n",solution.ByName_("L2"))
}

/*aprox := m.ProveAll(`aprox(1.14,Y).`)
for _, solution := range aprox {
    //fmt.Printf("%s -> %s -> %s -> %s\n", solution.ByName_("X"), solution.ByName_("Y"), solution.ByName_("N"), solution.ByName_("Z"))
    fmt.Printf("%s",solution.ByName_("X"))
}*/

/*contratos := m.ProveAll(`valor_contrato(X,Y).`)
for _, solution := range contratos {
    fmt.Printf("%s contrato -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}
pagos_brutos := m.ProveAll(`valor_pago(X,Y).`)
for _, solution := range pagos_brutos {
    fmt.Printf("%s pago -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}
pagos_salud := m.ProveAll(`valor_descuento_salud(X,Y).`)
for _, solution := range pagos_salud {
    fmt.Printf("%s pago salud -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}
pagos_pension := m.ProveAll(`valor_descuento_pension(X,Y).`)
for _, solution := range pagos_pension {
    fmt.Printf("%s pago pension -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}
pagos_bruto := m.ProveAll(`valor_pago_bruto(X,Y).`)
for _, solution := range pagos_bruto {
    fmt.Printf("%s pago bruto -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
}*/

/*pagos_finales := m.ProveAll(`valores_finales(X,Y,2016).`)
for _, solution := range pagos_finales {
    fmt.Printf("%s lista -> %s \n", solution.ByName_("X"), solution.ByName_("Y"))
    r := strings.NewReplacer("(", "",")", "","[]", "","'.'", "",)
    cad := r.Replace(solution.ByName_("Y").String())
    fmt.Println(cad)
    pair := strings.Split(cad, ",")
    fmt.Println(pair[1])
}*/

}

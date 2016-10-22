package routers

import (
	"github.com/astaxie/beego"
)

func init() {



	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Generar",
			`/Generar`,
			[]string{"get"},
			nil})

		beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
	  beego.ControllerComments{
	    "Liquidacion",
	    `/Liquidacion`,
	    []string{"post"},
	    nil})
		beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PruebasController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PruebasController"],
	  beego.ControllerComments{
	    "Preliquidacion",
	    `/preliquidacion`,
	    []string{"get"},
	    nil})

}

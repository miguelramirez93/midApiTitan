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

}

package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetalleNovedadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NominaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:PreliquidacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}

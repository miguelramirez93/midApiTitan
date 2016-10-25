package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:DescuentosController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

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

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:InformacionProveedorController"],
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

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiTitan/controllers:NovedadAplicadaController"],
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

package controllers

import (
	"encoding/json"
	models "github.com/miguelramirez93/midApiTitan/models"
	"github.com/miguelramirez93/midApiTitan/golog"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/jung-kurt/gofpdf"
	//"fmt"
)

// oprations for Preliquidacion
type PreliquidacionController struct {
	beego.Controller
}

func (c *PreliquidacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Generar", c.Generar)
	c.Mapping("Liquidacion", c.Liquidacion)
}


//generar la preliquidacion
func (this *PreliquidacionController) Generar() {
	var postnomina string = ""
	var postdominio string = ""
	var preliquidacion string = ""
	if tnomina  := this.GetString("tnomina"); tnomina != "" {
			postnomina = postnomina +"&query=TipoContrato.Id:"+tnomina
	}
	if tdominio  := this.GetString("tdominio"); tdominio != "" {
			postdominio = postdominio +"&query=Dominio.Id:"+tdominio
	}
	if tpreliquidacion  := this.GetString("preliquidacion"); tpreliquidacion != "" {
			preliquidacion = tpreliquidacion
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "falta id de la preliquidacion"}
		this.ServeJSON()
	}
	var v []models.Predicado
	var datos_contrato []models.ContratoGeneral
	var datos_novedades []models.DetalleNovedad
	var predicados []models.Predicado
	var idDetaPre int
	var res interface{}
	if err := getJson("http://"+beego.AppConfig.String("Urlruler")+":"+beego.AppConfig.String("Portruler")+"/"+beego.AppConfig.String("Nsruler")+"/predicado?limit=0"+postdominio, &v); err == nil {
		//Tomar del json el nombre de la regla y guardarlo en arregloReglas

		if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general?limit=0"+postnomina, &datos_contrato); err == nil {
			var reglas string = ""
			var reglasbase string = ""
			var reglasinyectadas string = ""
			var arregloReglas = make([]string, len(v))

			//var respuesta []models.FormatoPreliqu
			for i := 0; i < len(v); i++ {
				arregloReglas[i] = v[i].Nombre
			}

			for i := 0; i < len(arregloReglas); i++ {
				reglasbase = reglasbase + arregloReglas[i]
			}

			for i := 0; i < len(datos_contrato); i++ {
				//solicitud de informacion de novedades de cada empleado si esta activa la novedad
				if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_novedad?limit=0&query=Estado:Activo,Persona:"+strconv.FormatInt(datos_contrato[i].Contratista.NumDocumento,10), &datos_novedades); err == nil {
						if(datos_novedades != nil){
								predicados = append(predicados,models.Predicado{Nombre:"factor('"+datos_contrato[i].Contratista.NomProveedor+"',"+"descuento,"+datos_novedades[0].Novedad.TipoNovedad+",'"+datos_novedades[0].Novedad.Nombre+"',"+strconv.FormatFloat(datos_novedades[0].Valor, 'f', -1, 64)+","+strconv.Itoa(datos_novedades[0].Vigencia)+")."} )
							} //regla de descuentos
				}

				a,m,_,_,_,_ := diff(datos_contrato[i].FechaInicio,datos_contrato[i].FechaFinal)
				meses_contrato := (a*12)+m

				predicados = append(predicados,models.Predicado{Nombre:"valor_contrato('"+datos_contrato[i].Contratista.NomProveedor+"',"+datos_contrato[i].ValorContrato+")."} )
				predicados = append(predicados,models.Predicado{Nombre:"duracion_contrato('"+datos_contrato[i].Contratista.NomProveedor+"',"+strconv.Itoa(meses_contrato)+",2016)."} )
				var arregloReglasInyectadas = make([]string, len(predicados))
				for i := 0; i < len(predicados); i++ {
					arregloReglasInyectadas[i] = predicados[i].Nombre
				}
				for i := 0; i < len(arregloReglasInyectadas); i++ {
					reglasinyectadas = reglasinyectadas + arregloReglasInyectadas[i]
				}
				reglas = reglasinyectadas+reglasbase
				//fmt.Print("Reglas: "+reglas)
				temp := golog.CargarReglas(reglas,"2016")
				Vneto := temp[0].Valor_neto

				//fmt.Print(" total: "+strconv.FormatFloat(datos_contrato[i].ValorContrato, 'f', 6, 64))
				Idpreliqu ,_ := strconv.Atoi(preliquidacion)
				pl :=  models.Preliquidacion{Id: Idpreliqu}
				persona :=  models.InformacionProveedor{NumDocumento:datos_contrato[i].Contratista.NumDocumento}
				detallepreliqu := models.DetallePreliquidacion{Persona: &persona, Valor : Vneto, Preliquidacion : &pl }
				if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion","POST",&idDetaPre ,&detallepreliqu); err == nil {

				}else{
					beego.Debug("error: ", err)
					this.Data["json"] = map[string]interface{}{"Mensaje": "error al guaradar el detalle"}
					this.ServeJSON()
				}
				for _, descuentos := range *temp[0].Descuentos {
			    descuentos.DetallePreliquidacion = &models.DetallePreliquidacion{Id: idDetaPre}
					if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/descuentos","POST",&res ,descuentos); err == nil {

					}else{
						beego.Debug("error: ", err)
						this.Data["json"] = map[string]interface{}{"Mensaje": "error al guaradar descuentos"}
						this.ServeJSON()
					}
			  }
				if(datos_novedades != nil){
					for _, novedades_aplicadas := range datos_novedades {
						nov := models.NovedadAplicada{DetallePreliquidacion: &models.DetallePreliquidacion{Id: idDetaPre},DetalleNovedad: &novedades_aplicadas}
						if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/novedad_aplicada","POST",&res ,nov); err == nil {

						}else{
							beego.Debug("error: ", err)
							this.Data["json"] = map[string]interface{}{"Mensaje": "error al guaradar descuentos"}
							this.ServeJSON()
						}
					}
				}
				//respuesta = append(respuesta,models.FormatoPreliqu{Respuesta: &temp[0]} )
				predicados = nil;
				datos_novedades = nil
				reglasinyectadas = ""
			}
			this.Data["json"] = map[string]interface{}{"Mensaje": "Preliquidacion generada correctamente"}
			this.ServeJSON()
		} else {
			beego.Debug("error: ", err)
			this.Data["json"] = map[string]interface{}{"Mensaje": "no se pudo generar la preliquidacion api crud no encontrado"}
			this.ServeJSON()
		}

	} else {
		beego.Debug("error: ", err)
		this.Data["json"] = map[string]interface{}{"Mensaje": "no se pudo generar la preliquidacion api ruler no encontrado"}
		this.ServeJSON()
	}
}

// @Title Post
// @Description create Preliquidacion
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {int} models.Preliquidacion
// @Failure 403 body is empty
// @router / [post]
func (c *PreliquidacionController) Post() {
	var v models.Preliquidacion
	var respuesta map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/preliquidacion","POST",&respuesta ,&v)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Preliquidacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Preliquidacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PreliquidacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var v models.Preliquidacion
	var respuesta models.Preliquidacion
	err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/preliquidacion/"+idStr,"GET",&respuesta ,&v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = respuesta
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Preliquidacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Preliquidacion
// @Failure 403
// @router / [get]
func (c *PreliquidacionController) GetAll() {
	var poststr string = ""
	if v := c.GetString("fields"); v != "" {
		if(poststr == ""){
			poststr = poststr +"?fields="+v
		}else{
			poststr = poststr +"&fields="+v
		}

	}
	if v := c.GetString("limit"); v != "" {
		if(poststr == ""){
			poststr = poststr +"?limit="+v
		}else{
			poststr = poststr +"&limit="+v
		}

	}
	if v := c.GetString("offset"); v != ""{
		if(poststr == ""){
			poststr = poststr +"?offset="+v
		}else{
			poststr = poststr +"&offset="+v
		}
	}
	if v := c.GetString("sortby"); v != "" {
		if(poststr == ""){
			poststr = poststr +"?sortby="+v
		}else{
			poststr = poststr +"&sortby="+v
		}
	}
	if v := c.GetString("order"); v != "" {
		if(poststr == ""){
			poststr = poststr +"?order="+v
		}else{
			poststr = poststr +"&order="+v
		}

	}
	if v := c.GetString("query"); v != "" {
		if(poststr == ""){
			poststr = poststr +"?query="+v
		}else{
			poststr = poststr +"&query="+v
		}

	}
	var preliquidacion []models.Preliquidacion
	getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/preliquidacion"+poststr, &preliquidacion)
	c.Data["json"] = preliquidacion
	c.ServeJSON()
}

// @Title Update
// @Description update the Preliquidacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 200 {object} models.Preliquidacion
// @Failure 403 :id is not int
// @router /:id [put]
func (this *PreliquidacionController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	var data models.Preliquidacion
	var respuesta string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Debug("error: ", err)
		respuesta =  "No se recibieron los datos correctamente"
	}
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/preliquidacion/"+idStr,"PUT",&respuesta ,&data)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

// @Title Delete
// @Description delete the Preliquidacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *PreliquidacionController) Delete() {
	id := this.Ctx.Input.Param(":id")
	var respuesta string
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/preliquidacion/"+id,"DELETE",&respuesta ,nil)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

//funcion de generacion para pdf

func (this *PreliquidacionController) Liquidacion() {
	var data []models.DetallePreliquidacion
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf2 := gofpdf.New("P", "mm", "A4", "")
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Debug("error: ", err)
		respuesta :=  "No se recibieron los datos correctamente"
		this.Data["json"] = respuesta
		this.ServeJSON()
	}

	pdf.AddPage()
//	pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")

	pdf.SetFont("Arial", "B", 16)
/*
	for i:=0 ; i<4 ; i++ {
		pdf.CellFormat(120, 7, "dato", "1", 0, "", false, 0, "")
}
pdf.Ln(-1)
for i:=0 ;i<4 ;i++ {
		pdf.CellFormat(120, 6, "ioo", "1", 0, "", false, 0, "")
		pdf.CellFormat(120, 6, "iooi", "1", 0, "", false, 0, "")
		pdf.Ln(-1)
}*/
//	pdf.Image(example.ImageFile("logo.png"), 10, 10, 30, 0, false, "", 0, "")
//	pdf.Ln(20)

	pdf.Cell(70, 10, "DETALLE DE PAGOS")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 14)

	pdf.CellFormat(70, 10, "Nombre Proveeedor", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Persona.NomProveedor, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10, "Numero Documento", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,strconv.FormatInt(data[0].Persona.NumDocumento ,10), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Periodo de nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Periodo, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10, "Id preliquidacion", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Id), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Nombre preliquidacion", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nombre, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Id nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Nomina.Id), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Vinculacion", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Nomina.Vinculacion), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Nombre, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Descripcion nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Descripcion, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Tipo nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nomina.TipoNomina, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Estado de nomina", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Estado, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Id usuario", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,strconv.FormatInt(data[0].Preliquidacion.IdUsuario,10), "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Preliquidacion estado", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Estado, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10,"Preliquidacion descripcion", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Preliquidacion.Descripcion, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(70, 10, "Valor", "1", 0, "", false, 0, "")
	pdf.CellFormat(120, 10,data[0].Valor, "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	pdf2.AddPage()
pdf2.SetFont("Arial", "B", 16)
		pdf2.Cell(70, 10, "ORDEN DE PAGO")
		pdf2.Ln(20)
		pdf2.SetFont("Arial", "", 14)

		pdf2.CellFormat(70, 10, "Nombre Proveeedor", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Persona.NomProveedor, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Numero Documento", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.FormatInt(data[0].Persona.NumDocumento ,10), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Direccion", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Persona.Direccion, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Entidad bancaria", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.FormatFloat(data[0].Persona.IdEntidadBancaria,'g',1,64), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Tipo cuenta", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Persona.TipoCuentaBancaria, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Numero cuenta", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Persona.NumCuentaBancaria, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Periodo de nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Periodo, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Id preliquidacion", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Id), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Nombre preliquidacion", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nombre, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Id nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Nomina.Id), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Vinculacion", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.Itoa(data[0].Preliquidacion.Nomina.Vinculacion), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Nombre, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Descripcion nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Descripcion, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Tipo nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nomina.TipoNomina, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Estado de nomina", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Nomina.Estado, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Id usuario", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,strconv.FormatInt(data[0].Preliquidacion.IdUsuario,10), "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Preliquidacion estado", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Estado, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10,"Preliquidacion descripcion", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Preliquidacion.Descripcion, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)

		pdf2.CellFormat(70, 10, "Valor", "1", 0, "", false, 0, "")
		pdf2.CellFormat(120, 10,data[0].Valor, "1", 0, "", false, 0, "")
		pdf2.Ln(-1)
//	pdf.Cell(0, 10,t.Format(data[0].Preliquidacion.Fecha)) //(time)
	//pdf.Cell(0, 10,Time.String(data[0].Preliquidacion.FechaFin))//(time)
	//pdf.Cell(0, 10,Time.String(data[0].Preliquidacion.FechaInicio)) //(time)
	pdf2.OutputFileAndClose("ordenpago.pdf")
	pdf.OutputFileAndClose("detalle.pdf")

}

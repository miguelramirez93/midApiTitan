package controllers

import (
	"encoding/json"

	models "github.com/miguelramirez93/midApiTitan/models"

"strconv"

	"github.com/astaxie/beego"
)

// oprations for DetallePreliquidacion
type DetallePreliquidacionController struct {
	beego.Controller
}

func (c *DetallePreliquidacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Informe", c.Informe)
}

// @Title Post
// @Description create DetallePreliquidacion
// @Param	body		body 	models.DetallePreliquidacion	true		"body for DetallePreliquidacion content"
// @Success 201 {int} models.DetallePreliquidacion
// @Failure 403 body is empty
// @router / [post]
func (this *DetallePreliquidacionController) Post() {
	var v models.DetallePreliquidacion
	var respuesta map[string]interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion","POST",&respuesta ,&v)
		this.Data["json"] = respuesta
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// @Title Get
// @Description get DetallePreliquidacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.DetallePreliquidacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DetallePreliquidacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var v models.DetallePreliquidacion
	var respuesta models.DetallePreliquidacion
	err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion/"+idStr,"GET",&respuesta ,&v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = respuesta
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get DetallePreliquidacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.DetallePreliquidacion
// @Failure 403
// @router / [get]
func (c *DetallePreliquidacionController) GetAll() {
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
	var detprel []models.DetallePreliquidacion
	getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion"+poststr, &detprel)
	c.Data["json"] = detprel
	c.ServeJSON()
}

// @Title Update
// @Description update the DetallePreliquidacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.DetallePreliquidacion	true		"body for DetallePreliquidacion content"
// @Success 200 {object} models.DetallePreliquidacion
// @Failure 403 :id is not int
// @router /:id [put]
func (this *DetallePreliquidacionController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	var data models.DetallePreliquidacion
	var respuesta string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Debug("error: ", err)
		respuesta = "No se recibieron los datos correctamente"
	}
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion/"+idStr,"PUT",&respuesta ,&data)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

// @Title Delete
// @Description delete the DetallePreliquidacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *DetallePreliquidacionController) Delete() {
	id := this.Ctx.Input.Param(":id")
	var respuesta string
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion/"+id,"DELETE",&respuesta ,nil)
	this.Data["json"] = respuesta
	this.ServeJSON()
}
func (this *DetallePreliquidacionController) Informe() {
	id_preliqu := this.GetString("id")
	var detalle_preliquidacion []models.DetallePreliquidacion
	err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion?limit=0&query=Preliquidacion.Id:"+id_preliqu,"GET",&detalle_preliquidacion ,nil)

	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		var respuesta = make([]models.Respuesta, len(detalle_preliquidacion))
		for i := 0; i < len(respuesta); i++ {

			var descuentos []models.Descuentos
			id_detalle := strconv.Itoa(detalle_preliquidacion[i].Id)
			err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/descuentos?query=DetallePreliquidacion.Id:"+id_detalle,"GET",&descuentos ,nil)

			if err != nil {
				this.Data["json"] = err.Error()
			} else {

			}
			var novedades []models.DetalleNovedad
			err2 := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/novedades&query=DetallePreliquidacion.Id:"+id_detalle,"GET",&novedades ,nil)

			if err2 != nil {
				this.Data["json"] = err2.Error()
			} else {

			}
			respuesta[i].NumDocumento = detalle_preliquidacion[i].Persona.NumDocumento
			respuesta[i].Nombre_Cont= detalle_preliquidacion[i].Persona.NomProveedor
			respuesta[i].Valor_bruto= detalle_preliquidacion[i].ValorBruto
			respuesta[i].Valor_neto= detalle_preliquidacion[i].Valor
			respuesta[i].Descuentos= &descuentos
			respuesta[i].Novedades= &novedades
		}

		this.Data["json"] = respuesta
		this.ServeJSON()
	}

}

package controllers

import (
	"encoding/json"
	models "github.com/miguelramirez93/midApiTitan/models"
	"github.com/astaxie/beego"
)

// oprations for ContratoGeneral
type ContratoGeneralController struct {
	beego.Controller
}

func (c *ContratoGeneralController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create ContratoGeneral
// @Param	body		body 	models.ContratoGeneral	true		"body for ContratoGeneral content"
// @Success 201 {int} models.ContratoGeneral
// @Failure 403 body is empty
// @router / [post]
func (this *ContratoGeneralController) Post() {
	var v models.ContratoGeneral
	var respuesta map[string]interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &v); err == nil {
		sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general","POST",&respuesta ,&v)
		this.Data["json"] = respuesta
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// @Title Get
// @Description get ContratoGeneral by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ContratoGeneral
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ContratoGeneralController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var v models.ContratoGeneral
	var respuesta models.ContratoGeneral
	err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general/"+idStr,"GET",&respuesta ,&v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = respuesta
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get ContratoGeneral
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ContratoGeneral
// @Failure 403
// @router / [get]
func (c *ContratoGeneralController) GetAll() {
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
	var detprel []models.ContratoGeneral
	getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general"+poststr, &detprel)
	c.Data["json"] = detprel
	c.ServeJSON()
}

// @Title Update
// @Description update the ContratoGeneral
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ContratoGeneral	true		"body for ContratoGeneral content"
// @Success 200 {object} models.ContratoGeneral
// @Failure 403 :id is not int
// @router /:id [put]
func (this *ContratoGeneralController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	var data models.ContratoGeneral
	var respuesta string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Debug("error: ", err)
		respuesta = "No se recibieron los datos correctamente"
	}
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general/"+idStr,"PUT",&respuesta ,&data)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

// @Title Delete
// @Description delete the ContratoGeneral
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ContratoGeneralController) Delete() {
	id := this.Ctx.Input.Param(":id")
	var respuesta string
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general/"+id,"DELETE",&respuesta ,nil)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

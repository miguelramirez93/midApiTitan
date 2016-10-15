package controllers

import (
	"encoding/json"
	 models "github.com/miguelramirez93/midApiTitan/models"
	"github.com/astaxie/beego"
)

// oprations for InformacionProveedor
type InformacionProveedorController struct {
	beego.Controller
}

func (c *InformacionProveedorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create InformacionProveedor
// @Param	body		body 	models.InformacionProveedor	true		"body for InformacionProveedor content"
// @Success 201 {int} models.InformacionProveedor
// @Failure 403 body is empty
// @router / [post]
func (c *InformacionProveedorController) Post() {
	var v models.InformacionProveedor
	var respuesta map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_proveedor","POST",&respuesta ,&v)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get InformacionProveedor by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.InformacionProveedor
// @Failure 403 :id is empty
// @router /:id [get]
func (c *InformacionProveedorController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var v models.InformacionProveedor
	var respuesta models.InformacionProveedor
	err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_proveedor/"+idStr,"GET",&respuesta ,&v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = respuesta
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get InformacionProveedor
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.InformacionProveedor
// @Failure 403
// @router / [get]
func (c *InformacionProveedorController) GetAll() {
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
	var proveedores []models.InformacionProveedor
	getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_proveedor"+poststr, &proveedores)
	c.Data["json"] = proveedores
	c.ServeJSON()
}

// @Title Update
// @Description update the InformacionProveedor
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.InformacionProveedor	true		"body for InformacionProveedor content"
// @Success 200 {object} models.InformacionProveedor
// @Failure 403 :id is not int
// @router /:id [put]
func (this *InformacionProveedorController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	var data models.InformacionProveedor
	var respuesta string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		beego.Debug("error: ", err)
		respuesta = "No se recibieron los datos correctamente"
		this.Data["json"] = respuesta
		this.ServeJSON()
	}
	err2 := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_proveedor/"+idStr,"PUT",&respuesta ,&data)
	if(err2 != nil){
		beego.Debug("error: ", err2)
	}
	this.Data["json"] = respuesta
	this.ServeJSON()
}

// @Title Delete
// @Description delete the InformacionProveedor
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *InformacionProveedorController) Delete() {
	id := this.Ctx.Input.Param(":id")
	var respuesta string
	sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/informacion_proveedor/"+id,"DELETE",&respuesta ,nil)
	this.Data["json"] = respuesta
	this.ServeJSON()
}

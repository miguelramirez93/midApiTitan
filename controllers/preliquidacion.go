package controllers

import (
	"encoding/json"

	models "github.com/miguelramirez93/midApiTitan/models"
	"github.com/miguelramirez93/midApiTitan/golog"

	"github.com/astaxie/beego"
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
}


//generar la preliquidacion
func (this *PreliquidacionController) Generar() {
	var v []models.Predicado
	var datos_contrato []models.ContratoGeneral
	var predicados []models.Predicado
	if err := getJson("http://"+beego.AppConfig.String("Urlruler")+":"+beego.AppConfig.String("Portruler")+"/"+beego.AppConfig.String("Nsruler")+"/predicado?limit=0", &v); err == nil {
		//Tomar del json el nombre de la regla y guardarlo en arregloReglas

		if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/contrato_general?limit=0", &datos_contrato); err == nil {
			var reglas string = ""
			var reglasbase string = ""
			var reglasinyectadas string = ""
			var arregloReglas = make([]string, len(v))

			var respuesta []models.FormatoPreliqu
			for i := 0; i < len(v); i++ {
				arregloReglas[i] = v[i].Nombre
			}

			for i := 0; i < len(arregloReglas); i++ {
				reglasbase = reglasbase + arregloReglas[i]
			}

			for i := 0; i < len(datos_contrato); i++ {
				predicados = append(predicados,models.Predicado{Nombre:"categoria('"+datos_contrato[i].Contratista.NomProveedor+"',"+"asociado)."} )
				predicados = append(predicados,models.Predicado{Nombre:"vinculacion('"+datos_contrato[i].Contratista.NomProveedor+"',"+"hc)."} )
				predicados = append(predicados,models.Predicado{Nombre:"horas('"+datos_contrato[i].Contratista.NomProveedor+"',"+"4500)."} )
				predicados = append(predicados,models.Predicado{Nombre:"duracion_contrato('"+datos_contrato[i].Contratista.NomProveedor+"',"+"1)."} )
				var arregloReglasInyectadas = make([]string, len(predicados))
				for i := 0; i < len(predicados); i++ {
					arregloReglasInyectadas[i] = predicados[i].Nombre
				}
				for i := 0; i < len(arregloReglasInyectadas); i++ {
					reglasinyectadas = reglasinyectadas + arregloReglasInyectadas[i]
				}
				reglas = reglasinyectadas+reglasbase
				temp := golog.CargarReglas(reglas)
				respuesta = append(respuesta,models.FormatoPreliqu{Contrato: &datos_contrato[i], Respuesta: &temp[1]} )
				predicados = nil;
				reglasinyectadas = ""
			}
			this.Data["json"] = respuesta
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

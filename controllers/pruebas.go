package controllers

import (

	models "github.com/miguelramirez93/midApiTitan/models"
	"github.com/miguelramirez93/midApiTitan/golog"
	"strconv"
	"github.com/astaxie/beego"

)

// oprations for Preliquidacion
type PruebasController struct {
	beego.Controller
}

func (c *PruebasController) URLMapping() {
	c.Mapping("Preliquidacion", c.Preliquidacion)
}

func (this* PruebasController) Preliquidacion(){
	var postdominio string = ""
	var preliquidacion string = ""
	var valor_contrato string = ""
	var duracion_contrato string = ""
	var doc_contratista string = ""
	var valor_novedad string = "0"
	var t_descuento_novedad string = ""
	if tdominio  := this.GetString("tdominio"); tdominio != "" {
			postdominio = postdominio +"&query=Dominio.Id:"+tdominio
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "no se definio el dominio de las reglas"}
		this.ServeJSON()
	}
	if tpreliquidacion  := this.GetString("preliquidacion"); tpreliquidacion != "" {
			preliquidacion = tpreliquidacion
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "falta id de la preliquidacion"}
		this.ServeJSON()
	}
	if vcontrato  := this.GetString("vcontrato"); vcontrato != "" {
		valor_contrato = vcontrato
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "no se definio el valor del contrato"}
		this.ServeJSON()
	}
	if dcontrato  := this.GetString("dcontrato"); dcontrato != "" {
		duracion_contrato = dcontrato
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "no se definio duracion del contrato"}
		this.ServeJSON()
	}
	if ndocumento  := this.GetString("ndocumento"); ndocumento != "" {
		doc_contratista = ndocumento
	}else{
		this.Data["json"] = map[string]interface{}{"Mensaje": "no se definio el documento del empleado"}
		this.ServeJSON()
	}
	if vnovedad  := this.GetString("vnovedad"); vnovedad != "" {
		valor_novedad = vnovedad
		if tdescuento  := this.GetString("tdescuento"); tdescuento != "" {
			t_descuento_novedad = tdescuento
		}else{
			this.Data["json"] = map[string]interface{}{"Mensaje": "no se definio el tipo de descuento de la novedad"}
			this.ServeJSON()
		}
	}else{


	}

	var v []models.Predicado
	var predicados []models.Predicado
	var reglas string = ""
	var reglasbase string = ""
	var reglasinyectadas string = ""

if err := getJson("http://"+beego.AppConfig.String("Urlruler")+":"+beego.AppConfig.String("Portruler")+"/"+beego.AppConfig.String("Nsruler")+"/predicado?limit=0"+postdominio, &v); err == nil {
	var arregloReglas = make([]string, len(v))
	for i := 0; i < len(v); i++ {
		arregloReglas[i] = v[i].Nombre
	}

	for i := 0; i < len(arregloReglas); i++ {
		reglasbase = reglasbase + arregloReglas[i]
	}
	predicados = append(predicados,models.Predicado{Nombre:"valor_contrato('"+doc_contratista+"',"+valor_contrato+")."} )
	predicados = append(predicados,models.Predicado{Nombre:"duracion_contrato('"+doc_contratista+"',"+duracion_contrato+",2016)."} )
	if(valor_novedad != "0"){
			predicados = append(predicados,models.Predicado{Nombre:"factor('"+doc_contratista+"',"+"descuento,"+t_descuento_novedad+",prueba,"+valor_novedad+",1)."})
	}
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
	num_doc, err := strconv.ParseInt(doc_contratista, 10, 64)
	preliqu, err := strconv.Atoi(preliquidacion)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"Mensaje": "error al convertir el dato"}
		this.ServeJSON()
	}
	pl :=  models.Preliquidacion{Id: preliqu}
	persona :=  models.InformacionProveedor{NumDocumento: num_doc}
	detallepreliqu := models.DetallePreliquidacion{Persona: &persona, Valor : Vneto, Preliquidacion : &pl }
	this.Data["json"] = detallepreliqu
	this.ServeJSON()
}




}

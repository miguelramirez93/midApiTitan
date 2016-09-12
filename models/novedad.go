package models

type Novedad struct {
	Id          int               `pk;orm:"column(id)"`
	IdCategoria *CategoriaNovedad `orm:"column(id_categoria);rel(fk)"`
	Estado      string            `orm:"column(estado)"`
	Nombre      string            `orm:"column(nombre)"`
	Simbolo     string            `orm:"column(simbolo)"`
	Naturaleza  string            `orm:"column(naturaleza)"`
	Descripcion string            `orm:"column(descripcion)"`
	TipoNovedad string            `orm:"column(tipo_novedad)"`
	Formula     string            `orm:"column(formula)"`
	IdProveedor *InformacionProveedor           `orm:"column(id_proveedor);rel(fk)"`
}

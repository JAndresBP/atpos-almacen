package producto_pkg

type Producto struct {
	Id             int    `bson:"id"`
	Nombre         string `bson:"nombre"`
	PrecioUnitario int    `bson:"precioUnitario"`
}

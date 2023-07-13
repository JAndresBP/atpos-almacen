package venta_pkg

import (
	"atpos-almacen-api/dominio/producto_pkg"
)

type Estado int

const (
	InProgress Estado = iota
	Completed
)

type ItemVenta struct {
	Producto producto_pkg.Producto `bson:"producto"`
	Cantidad int                   `bson:"cantidad"`
}

type Venta struct {
	Id        int               `bson:"id"`
	Productos map[int]ItemVenta `bson:"productos"`
	Total     int               `bson:"total"`
	Estado    Estado            `bson:"estado"`
	Timestamp int64				`bson:"timestamp"`
}

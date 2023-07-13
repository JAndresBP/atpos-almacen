package dominio

import (
	"atpos-almacen-api/dominio/producto_pkg"
	"atpos-almacen-api/dominio/venta_pkg"
)

type Entity interface {
	venta_pkg.Venta | producto_pkg.Producto
}

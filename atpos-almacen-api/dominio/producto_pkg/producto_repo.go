package producto_pkg

type IProductoRepo interface {
	ObtenerProducto(id int) Producto
}

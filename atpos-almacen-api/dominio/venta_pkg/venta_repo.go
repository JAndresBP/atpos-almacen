package venta_pkg

type IVentaRepo interface {
	ObtenerVenta(id int) (Venta, bool)
	CrearVenta() Venta
	ActualizarVenta(venta Venta) Venta
}

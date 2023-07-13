package aplicacion

import (
	"atpos-almacen-api/dominio"
	"atpos-almacen-api/dominio/producto_pkg"
	"atpos-almacen-api/dominio/venta_pkg"
	"encoding/json"
	"net/http"
	"time"
)

type ProcesadorVenta struct {
	VR venta_pkg.IVentaRepo
	PR producto_pkg.IProductoRepo
	SS dominio.ISyncService
}

type RegistrarProducto struct {
	IdVenta    int `json:"idVenta"`
	IdProducto int `json:"idProducto"`
}

type RegistrarYCompletar struct {
	IdProductos []int `json:"productos"`
}

func (pv *ProcesadorVenta) RegistrarProducto(request *RegistrarProducto) (venta_pkg.Venta, bool) {

	var idVenta = request.IdVenta
	var venta venta_pkg.Venta
	ok := false
	if idVenta > 0 {
		venta, ok = pv.VR.ObtenerVenta(idVenta)
	}

	if !ok {
		venta = pv.VR.CrearVenta()
	}

	if venta.Estado == venta_pkg.Completed {
		return venta, false
	}

	producto := pv.PR.ObtenerProducto(request.IdProducto)

	itemVenta, ok := venta.Productos[producto.Id]
	if ok {
		itemVenta.Cantidad += 1
		venta.Productos[producto.Id] = itemVenta
	} else {
		venta.Productos[producto.Id] = venta_pkg.ItemVenta{Producto: producto, Cantidad: 1}
	}

	pv.VR.ActualizarVenta(venta)
	return venta, true
}

func (pv *ProcesadorVenta) CompletarVenta(idVenta int) (venta_pkg.Venta, int) {
	venta, ok := pv.VR.ObtenerVenta(idVenta)
	if !ok {
		return venta, http.StatusNotFound
	}

	if venta.Estado == venta_pkg.Completed {
		return venta, http.StatusNotModified
	}

	venta.Estado = venta_pkg.Completed
	venta.Timestamp = time.Now().UnixMilli()
	data, _ := json.Marshal(venta)
	pv.SS.Publish(data)
	pv.VR.ActualizarVenta(venta)
	return venta, http.StatusOK
}

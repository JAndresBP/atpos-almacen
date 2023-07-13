package repositorios

import "atpos-almacen-api/dominio/producto_pkg"

type ProductoRepo struct {
	cache map[int]producto_pkg.Producto
}

var productoRepoInstance *ProductoRepo

func GetProductoRepo(DB *MongoDbCtx) *ProductoRepo {
	if productoRepoInstance == nil {
		productoRepoInstance = &ProductoRepo{
			cache: map[int]producto_pkg.Producto{
				1:  {Id: 1, Nombre: "producto A", PrecioUnitario: 1000},
				2:  {Id: 2, Nombre: "producto B", PrecioUnitario: 2000},
				3:  {Id: 3, Nombre: "producto C", PrecioUnitario: 3000},
				4:  {Id: 4, Nombre: "producto D", PrecioUnitario: 4000},
				5:  {Id: 5, Nombre: "producto E", PrecioUnitario: 5000},
				6:  {Id: 6, Nombre: "producto F", PrecioUnitario: 6000},
				7:  {Id: 7, Nombre: "producto G", PrecioUnitario: 7000},
				8:  {Id: 8, Nombre: "producto H", PrecioUnitario: 8000},
				9:  {Id: 9, Nombre: "producto I", PrecioUnitario: 9000},
				10: {Id: 10, Nombre: "producto J", PrecioUnitario: 10000},
			},
		}
	}

	return productoRepoInstance
}

func (pr *ProductoRepo) ObtenerProducto(id int) producto_pkg.Producto {
	return pr.cache[id]
}

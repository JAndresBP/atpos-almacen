package main

import (
	"atpos-almacen-api/aplicacion"
	"atpos-almacen-api/dominio"
	"atpos-almacen-api/dominio/producto_pkg"
	"atpos-almacen-api/dominio/venta_pkg"
	"atpos-almacen-api/infrastructura"
	"atpos-almacen-api/infrastructura/repositorios"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type services struct {
	VR venta_pkg.IVentaRepo
	PR producto_pkg.IProductoRepo
	SS dominio.ISyncService
	DB repositorios.MongoDbCtx
}

var serviceCollection *services

func RegistrarProducto(c *gin.Context) {
	procesadorVenta := aplicacion.ProcesadorVenta{VR: serviceCollection.VR, PR: serviceCollection.PR}

	var request aplicacion.RegistrarProducto
	c.BindJSON(&request)
	venta, ok := procesadorVenta.RegistrarProducto(&request)
	if ok {
		c.IndentedJSON(http.StatusOK, venta)
	} else {
		c.AbortWithStatus(http.StatusConflict)
	}
}

func CompletarVenta(c *gin.Context) {
	procesadorVenta := aplicacion.ProcesadorVenta{VR: serviceCollection.VR, PR: serviceCollection.PR, SS: serviceCollection.SS}
	idVenta, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	venta, status := procesadorVenta.CompletarVenta(int(idVenta))
	c.IndentedJSON(status, venta)
}

func RegistrarYCompletar(c *gin.Context) {
	procesadorVenta := aplicacion.ProcesadorVenta{VR: serviceCollection.VR, PR: serviceCollection.PR, SS: serviceCollection.SS}

	var request aplicacion.RegistrarProducto
	c.BindJSON(&request)
	venta, ok := procesadorVenta.RegistrarProducto(&request)
	if ok {
		_, status := procesadorVenta.CompletarVenta(venta.Id)
		c.IndentedJSON(status, venta)
	} else {
		c.AbortWithStatus(http.StatusConflict)
	}
}

func main() {
	router := gin.Default()
	mongoserver := os.Getenv("MONGO_SERVER")
	rabbitmqserver := os.Getenv("RABBITMQ_SERVER")
	DB := repositorios.GetMongoDbCtx("mongodb://root:example@" + mongoserver)

	serviceCollection = &services{
		VR: repositorios.GetVentaRepo(DB),
		PR: repositorios.GetProductoRepo(DB),
		SS: infrastructura.GetSyncService("amqp://guest:guest@" + rabbitmqserver),
	}

	defer serviceCollection.SS.Close()
	defer DB.Close()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, mongoserver)
	})
	router.POST("/registrar", RegistrarProducto)
	router.POST("/checkout/:id", CompletarVenta)
	router.POST("/registerandcheck", RegistrarYCompletar)
	router.Run("0.0.0.0:3000")
}

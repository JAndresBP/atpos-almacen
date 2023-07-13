package repositorios

import (
	"atpos-almacen-api/dominio/venta_pkg"
	"atpos-almacen-api/infrastructura"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VentaRepo struct {
	cache      *infrastructura.LocalCache[int, venta_pkg.Venta]
	currentId  int
	DB         *MongoDbCtx
	Collection *mongo.Collection
}

var ventaRepoInstance *VentaRepo

func GetVentaRepo(DB *MongoDbCtx) *VentaRepo {
	if ventaRepoInstance == nil {

		ventaRepoInstance = &VentaRepo{
			cache:      infrastructura.NewLocalCache[int, venta_pkg.Venta](time.Duration(10) * time.Second),
			currentId:  0,
			DB:         DB,
			Collection: DB.Client.Database("almacen").Collection("ventas"),
		}
	}

	return ventaRepoInstance
}

func (vr *VentaRepo) ObtenerVenta(id int) (venta_pkg.Venta, bool) {
	venta, ok := vr.cache.Read(id)
	if !ok {
		filter := bson.D{primitive.E{Key: "id", Value: id}}
		vr.Collection.FindOne(vr.DB.Ctx, filter).Decode(venta)
	}
	return venta, ok
}

func (vr *VentaRepo) CrearVenta() venta_pkg.Venta {

	var timestamp int64 = time.Now().UTC().Add(time.Duration(24) * time.Hour).UnixMilli()
	vr.currentId += 1
	venta := venta_pkg.Venta{
		Id:        vr.currentId,
		Productos: make(map[int]venta_pkg.ItemVenta),
	}
	vr.cache.Update(venta.Id, venta, timestamp)
	return venta
}

func (vr *VentaRepo) ActualizarVenta(venta venta_pkg.Venta) venta_pkg.Venta {
	var timestamp int64 = time.Now().UTC().Add(time.Duration(24) * time.Hour).UnixMilli()
	vr.cache.Update(venta.Id, venta, timestamp)
	filter := bson.D{primitive.E{Key: "id", Value: venta.Id}}
	_, err := vr.Collection.ReplaceOne(context.TODO(), filter, venta, options.Replace().SetUpsert(true))
	if err != nil {
		log.Printf("%s", err)
	}
	return venta
}

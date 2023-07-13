using atpos_facturacion.Domain.ProductsAggregate;
using MongoDB.Bson.Serialization.Attributes;
namespace atpos_facturacion.Domain.SalesAggregate;

public class ProductItem {
    public Product? Producto {get;set;}
    public int Cantidad {get;set;}
}
public class Sale {
    public int Id {get;set;}
    public Dictionary<string,ProductItem>? Productos {get;set;}
    public int Total {get;set;}
    public int Estado {get;set;}
    public long Timestamp {get;set;}
    public long CentralTimestamp {get;set;}
}
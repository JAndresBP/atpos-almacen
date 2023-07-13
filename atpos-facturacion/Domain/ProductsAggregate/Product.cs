namespace atpos_facturacion.Domain.ProductsAggregate;
using MongoDB.Bson.Serialization.Attributes;
public class Product {
    public int Id {get;set;}
    public string? Nombre {get;set;}
    public int PrecioUnitario {get;set;}
}
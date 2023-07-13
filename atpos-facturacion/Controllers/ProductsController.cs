using Microsoft.AspNetCore.Mvc;
using atpos_facturacion.Domain.ProductsAggregate;
using MongoDB.Driver;

namespace atpos_facturacion.Controllers;

[ApiController]
[Route("[controller]")]
public class ProductsController : ControllerBase
{
    IMongoCollection<Product> _collection;

    public ProductsController(IMongoCollection<Product> collection)
    {
        _collection = collection;
    }

    [HttpGet]
    public async Task<ActionResult<IList<Product>>> Get()
    {
        var filter = Builders<Product>.Filter.Empty;
        var result = await _collection.Find(filter).ToListAsync();
        return Ok(result);
    }
}

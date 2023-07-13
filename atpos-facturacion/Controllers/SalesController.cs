using Microsoft.AspNetCore.Mvc;
using atpos_facturacion.Domain.SalesAggregate;
using MongoDB.Driver;

namespace atpos_facturacion.Controllers;

[ApiController]
[Route("[controller]")]
public class SalesController: ControllerBase{
    
    IMongoCollection<Sale> _collection;

    public SalesController(IMongoCollection<Sale> collection)
    {
        _collection = collection;
    }

    [HttpPost]
    public async Task<ActionResult> Post(Sale sale){
        sale.CentralTimestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();
        var filter = Builders<Sale>.Filter.Eq(r => r.Id, sale.Id);
        await _collection.ReplaceOneAsync(filter, sale, new ReplaceOptions{IsUpsert = true});
        return Ok();
    }
}    
    

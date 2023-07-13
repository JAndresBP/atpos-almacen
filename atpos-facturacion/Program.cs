using MongoDB.Driver;
using MongoDB.Bson;
using atpos_facturacion.Domain.ProductsAggregate;
using atpos_facturacion.Domain.SalesAggregate;

var builder = WebApplication.CreateBuilder(args);

var connectionString = Environment.GetEnvironmentVariable("MONGODB_URI") ?? "0.0.0.0:27017";
// Add services to the container.
var client = new MongoClient($"mongodb://root:example@{connectionString}");
var db = client.GetDatabase("central");
builder.Services.AddSingleton<IMongoCollection<Product>>(db.GetCollection<Product>("productos"));
builder.Services.AddSingleton<IMongoCollection<Sale>>(db.GetCollection<Sale>("ventas"));

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseAuthorization();
app.MapControllers();

app.Run();

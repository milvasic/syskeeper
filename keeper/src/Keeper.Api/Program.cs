using OpenTelemetry.Logs;
using OpenTelemetry.Metrics;
using OpenTelemetry.Trace;

var builder = WebApplication.CreateBuilder(args);

// Razor Pages
builder.Services.AddRazorPages();

// Health checks
builder.Services.AddHealthChecks()
	.AddNpgSql(builder.Configuration.GetConnectionString("Postgres")!);

// OpenAPI / Swagger
builder.Services.AddOpenApi();

// OpenTelemetry
builder.Logging.AddOpenTelemetry(logging =>
{
	logging.IncludeFormattedMessage = true;
	logging.IncludeScopes = true;
});

builder.Services.AddOpenTelemetry()
	.WithMetrics(metrics =>
	{
		metrics.AddAspNetCoreInstrumentation();
		metrics.AddHttpClientInstrumentation();
	})
	.WithTracing(tracing =>
	{
		tracing.AddAspNetCoreInstrumentation();
		tracing.AddHttpClientInstrumentation();
	});

var app = builder.Build();

if (!app.Environment.IsDevelopment())
{
	app.UseExceptionHandler("/Error");
	app.UseHsts();
}

app.UseHttpsRedirection();
app.UseStaticFiles();
app.UseRouting();
app.UseAuthorization();

// Swagger UI in development
if (app.Environment.IsDevelopment())
{
	app.MapOpenApi();
	app.UseSwaggerUI(options =>
	{
		options.SwaggerEndpoint("/openapi/v1.json", "Keeper API v1");
	});
}

// Health check endpoint
app.MapHealthChecks("/health");

app.MapRazorPages();

app.Run();

// Make Program accessible for integration tests
public partial class Program { }

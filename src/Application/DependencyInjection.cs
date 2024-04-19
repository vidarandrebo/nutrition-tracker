using Microsoft.Extensions.DependencyInjection;

namespace NutritionTracker.Application;

public static class DependencyInjection
{
    public static IServiceCollection AddApplicationServices(this IServiceCollection services)
    {
        services.AddMediatR(cfg => { cfg.RegisterServicesFromAssembly(typeof(Application.AssemblyName).Assembly); });
        return services;
    }
}
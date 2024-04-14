using Microsoft.Extensions.DependencyInjection;

namespace Application;

public static class AddApplication
{
    public static IServiceCollection AddApplicationServices(this IServiceCollection services)
    {
        services.AddMediatR(cfg => { cfg.RegisterServicesFromAssembly(typeof(Application.AssemblyName).Assembly); });
        return services;
    }
}
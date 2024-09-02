using System;
using System.IO;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Diagnostics;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Infrastructure.Identity;
using NutritionTracker.Infrastructure.Interceptors;

namespace NutritionTracker.Infrastructure;

public static class DependencyInjection
{
    public static IServiceCollection AddDatabase(this IServiceCollection services,
        IConfiguration configuration,
        IWebHostEnvironment environment)
    {
        services.AddScoped<ISaveChangesInterceptor, DispatchDomainEventsInterceptor>();
        if (environment.IsProduction())
        {
            Console.WriteLine("Production");
            Console.WriteLine(configuration.GetValue<string>("Database:Server"));
            var dbConnectionString = $"User ID={configuration.GetValue<string>("Database:User")};" +
                                     $"Password={configuration.GetValue<string>("Database:Password")};" +
                                     $"Server={configuration.GetValue<string>("Database:Server")};" +
                                     $"Port={configuration.GetValue<string>("Database:Port")};" +
                                     $"Database={configuration.GetValue<string>("Database:Name")};";

            services.AddDbContext<ApplicationDbContext>((serviceProvider, options) =>
            {
                options.AddInterceptors(serviceProvider.GetServices<ISaveChangesInterceptor>());
                options.UseNpgsql(dbConnectionString,
                    postgresOptions => { postgresOptions.MigrationsAssembly("NutritionTracker.Migrations.Postgres"); });
            });
        }
        else
        {
            services.AddDbContext<ApplicationDbContext>((serviceProvider, options) =>
            {
                var folder = configuration.GetValue<string>("Database:Folder");
                var filename = configuration.GetValue<string>("Database:File");
                options.UseSqlite($"Data Source={Path.Join(folder, filename)}",
                    sqliteOptions => { sqliteOptions.MigrationsAssembly("NutritionTracker.Migrations.Sqlite"); });
                options.AddInterceptors(serviceProvider.GetServices<ISaveChangesInterceptor>());
            });
        }

        services.AddScoped<IApplicationDbContext>(provider => provider.GetRequiredService<ApplicationDbContext>());

        return services;
    }
}
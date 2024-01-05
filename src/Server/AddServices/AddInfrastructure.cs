using System;
using Application.Interfaces;
using Infrastructure;
using Infrastructure.Interceptors;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Diagnostics;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Migrations.Postgres;

namespace Server.AddServices;

public static class AddInfrastructure
{
    public static IServiceCollection AddInfrastructureServices(this IServiceCollection services,
        IConfiguration configuration,
        IWebHostEnvironment environment)
    {
        services.AddScoped<ISaveChangesInterceptor, DispatchDomainEventsInterceptor>();
        if (environment.IsProduction())
        {
            Console.WriteLine("Production");
            DotEnv.Load(".env");
            var dbConnectionString = $"User ID={Environment.GetEnvironmentVariable("DB_USER")};" +
                                     $"Password={Environment.GetEnvironmentVariable("DB_PASSWD")};" +
                                     $"Server={Environment.GetEnvironmentVariable("DB_SERVER")};" +
                                     $"Port={Environment.GetEnvironmentVariable("DB_PORT")};" +
                                     $"Database={Environment.GetEnvironmentVariable("DB_NAME")};";

            services.AddDbContext<ApplicationDbContext>((serviceProvider, options) =>
            {
                options.AddInterceptors(serviceProvider.GetServices<ISaveChangesInterceptor>());
                options.UseNpgsql(dbConnectionString, postgresOptions =>
                {
                    postgresOptions.MigrationsAssembly(Provider.Assembly);
                });
            });
        }
        else
        {
            services.AddDbContext<ApplicationDbContext>((serviceProvider, options) =>
            {
                var folder = configuration.GetValue<string>("Database:Folder");
                var filename = configuration.GetValue<string>("Database:File");
                options.UseSqlite($"Data Source=nutrition-tracker.db", sqliteOptions =>
                    {
                    sqliteOptions.MigrationsAssembly(Migrations.Sqlite.Provider.Assembly);
                });
                options.AddInterceptors(serviceProvider.GetServices<ISaveChangesInterceptor>());
            });
        }

        services.AddScoped<IApplicationDbContext>(provider => provider.GetRequiredService<ApplicationDbContext>());
        services.AddIdentityApiEndpoints<IdentityUser>()
            .AddEntityFrameworkStores<ApplicationDbContext>();

        return services;
    }
}
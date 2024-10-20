using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application;
using NutritionTracker.Infrastructure;
using NutritionTracker.Infrastructure.Identity;
using NutritionTracker.Web.Identity;
using NutritionTracker.Web.Pages;
using Serilog.Core;

namespace NutritionTracker.Web;

public class Program
{
    public static async Task Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);
        builder.Configuration.LoadEnvToConfiguration(".env");

        // Add services to the container.
        builder.Services.AddRazorComponents()
            .AddInteractiveServerComponents();

        builder.Services.AddApplicationServices();
        builder.Services.AddCascadingAuthenticationState();
        builder.Services.AddScoped<IdentityUserAccessor>();
        builder.Services.AddScoped<IdentityRedirectManager>();
        builder.Services.AddScoped<AuthenticationStateProvider, IdentityRevalidatingAuthenticationStateProvider>();

        builder.Services.AddAuthentication(options =>
            {
                options.DefaultScheme = IdentityConstants.ApplicationScheme;
                options.DefaultSignInScheme = IdentityConstants.ExternalScheme;
            })
            .AddIdentityCookies();

        //var connectionString = builder.Configuration.GetConnectionString("DefaultConnection") ??
        //throw new InvalidOperationException("Connection string 'DefaultConnection' not found.");


//        builder.Services.AddDbContext<ApplicationDbContext>(options =>
//            options.UseSqlServer(connectionString));

        builder.Services.AddDatabase(builder.Configuration, builder.Environment);
        builder.Services.AddDatabaseDeveloperPageExceptionFilter();


        builder.Services.AddIdentityCore<ApplicationUser>(options => options.SignIn.RequireConfirmedAccount = true)
            .AddEntityFrameworkStores<ApplicationDbContext>()
            .AddSignInManager()
            .AddDefaultTokenProviders();

        builder.Services.AddSingleton<IEmailSender<ApplicationUser>, IdentityNoOpEmailSender>();

        var app = builder.Build();
        await app.Services.ApplyMigrations(app.Environment);

        // Configure the HTTP request pipeline.
        if (app.Environment.IsDevelopment())
        {
            app.UseMigrationsEndPoint();
        }
        else
        {
            app.UseExceptionHandler("/Error");
        }

        app.UseStaticFiles();
        app.UseAntiforgery();

        var logger = app.Services.GetRequiredService<ILogger<Program>>();

        app.MapRazorComponents<App>()
            .AddInteractiveServerRenderMode();

        // Add additional endpoints required by the Identity /Account Razor components.
        app.MapAdditionalIdentityEndpoints();

        foreach (var item in app.Configuration.AsEnumerable())
        {
            logger.LogInformation("{key} - {value}", item.Key, item.Value);
        }

        app.Run();
    }
}
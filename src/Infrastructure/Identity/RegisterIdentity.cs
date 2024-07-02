using System;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.DependencyInjection;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Infrastructure.Identity;

public static class Register
{
    public static IServiceCollection RegisterIdentity(this IServiceCollection services)
    {
        services.AddIdentity<ApplicationUser, IdentityRole<Guid>>(
                options =>
                {
                    options.User.RequireUniqueEmail = true;
                }
            )
            .AddEntityFrameworkStores<ApplicationDbContext>()
            .AddUserManager<UserManager<ApplicationUser>>();
        services.AddTransient<IIdentityService, IdentityService>();
        return services;
    }
}
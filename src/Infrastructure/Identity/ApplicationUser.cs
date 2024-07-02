using System;
using Microsoft.AspNetCore.Identity;

namespace NutritionTracker.Infrastructure.Identity;

public class ApplicationUser : IdentityUser<Guid>
{
    public Guid AccountId { get; set; }
}
using System.Threading;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using NutritionTracker.Domain.Accounts;
using NutritionTracker.Domain.FoodItems;
using NutritionTracker.Domain.FoodItems.Entities;
using NutritionTracker.Domain.Meals;
using NutritionTracker.Domain.Recipes;
using NutritionTracker.Domain.Recipes.Entities;

namespace NutritionTracker.Application.Interfaces;

public interface IApplicationDbContext
{
    public DbSet<FoodItem> FoodItems { get; set; }
    public DbSet<Recipe> Recipes { get; set; }
    public DbSet<Meal> Meals { get; set; }
    public DbSet<Account> Accounts { get; set; }
    Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}
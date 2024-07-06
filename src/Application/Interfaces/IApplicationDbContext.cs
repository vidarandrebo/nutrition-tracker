using System.Threading;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using NutritionTracker.Domain.Accounts.Entities;
using NutritionTracker.Domain.Diary.Entities;
using NutritionTracker.Domain.FoodItems.Entities;
using NutritionTracker.Domain.Recipes.Entities;

namespace NutritionTracker.Application.Interfaces;

public interface IApplicationDbContext
{
    public DbSet<FoodItem> FoodItems { get; set; }
    public DbSet<Recipe> Recipes { get; set; }
    public DbSet<Day> Days { get; set; }
    public DbSet<Account> Accounts { get; set; }
    Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}
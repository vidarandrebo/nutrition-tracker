using System.Threading;
using System.Threading.Tasks;
using NutritionTracker.Domain.FoodItems;
using Microsoft.EntityFrameworkCore;

namespace NutritionTracker.Application.Interfaces;

public interface IApplicationDbContext
{
    public DbSet<FoodItem> FoodItems { get; set; }
    Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}
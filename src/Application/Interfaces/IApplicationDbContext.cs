using System.Threading;
using System.Threading.Tasks;
using Domain.FoodItems;
using Microsoft.EntityFrameworkCore;

namespace Application.Interfaces;

public interface IApplicationDbContext
{
    public DbSet<FoodItem> FoodItems { get; set; }
    Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}
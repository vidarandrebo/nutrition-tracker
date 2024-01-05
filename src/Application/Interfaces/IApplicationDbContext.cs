using System.Threading;
using System.Threading.Tasks;

namespace Application.Interfaces;

public interface IApplicationDbContext
{
    Task<int> SaveChangesAsync(CancellationToken cancellationToken);
}
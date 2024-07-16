using System;
using System.Threading;
using System.Threading.Tasks;
using MediatR;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Accounts.Entities;

namespace NutritionTracker.Application.Accounts;

public class AddAccount
{
    public record Request(Guid Id) : IRequest<Unit>;

    public class Handler : IRequestHandler<Request, Unit>
    {
        private readonly IApplicationDbContext _db;

        public Handler(IApplicationDbContext db)
        {
            _db = db;
        }

        public async Task<Unit> Handle(Request request, CancellationToken ct)
        {
            var account = new Account(request.Id);
            _db.Accounts.Add(account);
            await _db.SaveChangesAsync(ct);
            return Unit.Value;
        }
    }
}
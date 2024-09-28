using System;
using System.Threading;
using System.Threading.Tasks;
using MediatR;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Accounts.Entities;

namespace NutritionTracker.Application.Accounts;

public class AddAccount
{
    public record Request(Guid Id) : IRequest<Unit>;

    public class Handler : IRequestHandler<Request, Unit>
    {
        private readonly IApplicationDbContext _db;
        private readonly ILogger<AddAccount> _logger;

        public Handler(IApplicationDbContext db, ILogger<AddAccount> logger)
        {
            _db = db;
            _logger = logger;
        }

        public async Task<Unit> Handle(Request request, CancellationToken ct)
        {
            var account = new Account(request.Id);
            _db.Accounts.Add(account);
            await _db.SaveChangesAsync(ct);
            _logger.LogInformation("Account {AccountId} added", account.Id);
            return Unit.Value;
        }
    }
}
using System;
using System.Threading;
using System.Threading.Tasks;
using FluentResults;
using MediatR;
using Microsoft.EntityFrameworkCore;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Accounts.Dtos;

namespace NutritionTracker.Application.Accounts
{
    public class GetAccount
    {
        public record Request(Guid Id) : IRequest<Result<AccountDto>>;

        public class Handler : IRequestHandler<Request, Result<AccountDto>>
        {
            private readonly IApplicationDbContext _db;

            public Handler(IApplicationDbContext db)
            {
                _db = db;
            }

            public async Task<Result<AccountDto>> Handle(Request request, CancellationToken ct)
            {
                var account = await _db.Accounts
                    .AsNoTracking()
                    .Include(x => x.MacronutrientGoal)
                    .Include(x => x.MicronutrientGoals)
                    .FirstOrDefaultAsync(a => a.Id == request.Id, ct);
                if (account is null)
                {
                    return Result.Fail("account not found");
                }

                return Result.Ok(account.ToDto());
            }
        }
    }
}
using Application.Interfaces;
using Domain.FoodItems;
using FluentResults;
using MediatR;
using Microsoft.EntityFrameworkCore;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace Application.FoodItems;

public class GetFoodItems

{
    public record Request() : IRequest<Result<FoodItemDTO[]>>;

    public class Handler : IRequestHandler<Request, Result<FoodItemDTO[]>>
    {
        private readonly IApplicationDbContext _db;

        public Handler(IApplicationDbContext db)
        {
            _db = db;
        }

        public async Task<Result<FoodItemDTO[]>> Handle(Request request, CancellationToken cancellationToken)
        {
            var foodItems = await _db.FoodItems
                .Select(f => f.ToDTO())
                .ToArrayAsync(cancellationToken);
            return Result.Ok(foodItems);
        }
    }
}
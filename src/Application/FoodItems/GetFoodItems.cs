using NutritionTracker.Domain.FoodItems;
using FluentResults;
using MediatR;
using Microsoft.EntityFrameworkCore;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Application.FoodItems;

public class GetFoodItems

{
    public record Request() : IRequest<Result<FoodItemDto[]>>;

    public class Handler : IRequestHandler<Request, Result<FoodItemDto[]>>
    {
        private readonly IApplicationDbContext _db;

        public Handler(IApplicationDbContext db)
        {
            _db = db;
        }

        public async Task<Result<FoodItemDto[]>> Handle(Request request, CancellationToken cancellationToken)
        {
            var foodItems = await _db.FoodItems
                .Select(f => f.ToDTO())
                .ToArrayAsync(cancellationToken);
            return Result.Ok(foodItems);
        }
    }
}
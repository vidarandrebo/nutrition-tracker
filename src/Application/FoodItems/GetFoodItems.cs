using System;
using System.Threading;
using System.Threading.Tasks;
using Application.Interfaces;
using Domain.FoodItems;
using FluentResults;
using MediatR;
using Microsoft.EntityFrameworkCore;

namespace Application.FoodItems;

public class GetFoodItems

{
    public record Request() : IRequest<Result<FoodItem[]>>;

    public class Handler : IRequestHandler<Request, Result<FoodItem[]>>
    {
        private readonly IApplicationDbContext _db;

        public Handler(IApplicationDbContext db)
        {
            _db = db;
        }

        public async Task<Result<FoodItem[]>> Handle(Request request, CancellationToken cancellationToken)
        {
            var foodItems = await _db.FoodItems.ToArrayAsync(cancellationToken);
            return Result.Ok(foodItems);
        }
    }
}
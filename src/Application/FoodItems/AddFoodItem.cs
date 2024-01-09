using System;
using System.Threading;
using System.Threading.Tasks;
using Application.Interfaces;
using Domain;
using FluentResults;
using MediatR;

namespace Application.FoodItems;

public class AddFoodItem
{
    public record Request(FoodItemForm form, Guid ownerId) : IRequest<Result>;

    public class Handler : IRequestHandler<Request, Result>
    {
        private readonly IApplicationDbContext _db;

        public Handler(IApplicationDbContext db)
        {
            _db = db;
        }

        public async Task<Result> Handle(Request request, CancellationToken cancellationToken)
        {
            var foodForm = request.form;
            var nutritionalContent =
                new NutritionalContent(foodForm.Protein, foodForm.Carbohydrate, foodForm.Fat, foodForm.KCal);
            var foodItem = new FoodItem(foodForm.Brand, foodForm.ProductName, nutritionalContent, request.ownerId);
            _db.FoodItems.Add(foodItem);
            await _db.SaveChangesAsync(cancellationToken);
            return Result.Ok();
        }
    }
}
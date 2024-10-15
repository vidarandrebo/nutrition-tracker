using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Bogus;
using MediatR;
using Microsoft.Extensions.Configuration;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Accounts.Events;
using NutritionTracker.Domain.FoodItems.Entities;

namespace NutritionTracker.Application.FoodItems.EventHandlers;

public class AccountCreatedEventHandler : INotificationHandler<AccountCreatedEvent>
{
    private readonly IApplicationDbContext _dbContext;
    private readonly IConfiguration _cfg;

    public AccountCreatedEventHandler(IApplicationDbContext dbContext, IConfiguration cfg)
    {
        _dbContext = dbContext;
        _cfg = cfg;
    }

    public async Task Handle(AccountCreatedEvent notification, CancellationToken cancellationToken)
    {
        if (_cfg.GetValue<string>("ENVIRONMENT") == "Production")
        {
            return;
        }

        var foodItemFaker = new Faker<FoodItem>()
            .RuleFor(f => f.Brand, f => f.Company.CompanyName())
            .RuleFor(f => f.ProductName, f => string.Join(' ', f.Lorem.Words()));

        var nutritionFaker = new Faker<Macronutrients>()
            .RuleFor(f => f.Protein, f => f.Random.Double(0.0, 33.3))
            .RuleFor(f => f.Carbohydrate, f => f.Random.Double(0.0, 33.3))
            .RuleFor(f => f.Fat, f => f.Random.Double(0.0, 33.3));
        for (int i = 0; i < 10000; i++)
        {
            var foodItem = new FoodItem("", "", new Macronutrients(0.0, 0.0, 0.0, 0.0), notification.AccountId,
                new List<Micronutrient>());

            foodItemFaker.Populate(foodItem);

            nutritionFaker.Populate(foodItem.Macronutrients);

            foodItem.Macronutrients.KCal = 4.0 * foodItem.Macronutrients.Protein + 4.0 * foodItem.Macronutrients.Carbohydrate + 9.0 * foodItem.Macronutrients.Fat;

            foodItem.Micronutrients = new List<Micronutrient>();

            _dbContext.FoodItems.Add(foodItem);
        }

        await _dbContext.SaveChangesAsync(cancellationToken);
    }
}
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Application.FoodItems.Commands;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.FoodItems.Contracts;

namespace NutritionTracker.Application.Tests.FoodItems;

public class AddFoodItemTest
{
    private readonly IApplicationDbContext _db;

    public AddFoodItemTest()
    {
        _db = DatabaseHelper.NewContext();
    }

    [Fact]
    public async Task AddOneFoodItem()
    {
        var ownerId = Guid.NewGuid();
        var foodForm = new PostFoodItemRequest();
        foodForm.Brand = "Brand";
        foodForm.ProductName = "ProductName";
        foodForm.Protein = 20;
        foodForm.Carbohydrate = 25;
        foodForm.Fat = 30;
        foodForm.KCal = 450;
        foodForm.Unit = "grams";
        var request = new AddFoodItem.Request(foodForm, ownerId);
        var handler = new AddFoodItem.Handler(_db);
        using var ctSource = new CancellationTokenSource(1000);
        await handler.Handle(request, ctSource.Token);

        var foodItem = _db.FoodItems.FirstOrDefault();
        Assert.NotNull(foodItem);
        Assert.NotNull(foodItem.Macronutrients);
        Assert.Equal(25, foodItem.Macronutrients.Carbohydrate);
    }
}
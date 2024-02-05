using Application.FoodItems;
using Application.Interfaces;
using Domain;
using Domain.FoodItems;
using Microsoft.EntityFrameworkCore;

namespace Application.Tests.FoodItems;

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
        var foodForm = new FoodItemForm();
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
        Assert.NotNull(foodItem.NutritionalContent);
        Assert.Equal(25, foodItem.NutritionalContent.Carbohydrate);
    }
}
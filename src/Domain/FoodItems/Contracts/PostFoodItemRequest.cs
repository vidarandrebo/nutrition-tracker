using System.Collections.Generic;
using FluentValidation;

namespace NutritionTracker.Domain.FoodItems.Contracts;

/// <summary>
/// Class representing the request data for a POST to endpoint /api/fooditem
/// </summary>
public class PostFoodItemRequest
{
    public string Brand { get; set; } = "";
    public string ProductName { get; set; } = "";
    public double Protein { get; set; }
    public double Carbohydrate { get; set; }
    public double Fat { get; set; }
    public double KCal { get; set; }
    public string Unit { get; set; } = "";

    public override string ToString()
    {
        return $"Brand: {Brand}\n" +
               $"ProductName {ProductName}\n" +
               $"Protein {Protein}\n" +
               $"Carbohydrate {Carbohydrate}\n" +
               $"Fat {Fat}\n" +
               $"KCal {KCal}\n" +
               $"Unit {Unit}";
    }
}

public class FoodItemValidator : AbstractValidator<PostFoodItemRequest>
{
    public FoodItemValidator()
    {
        var allowedUnits = new List<string>() { "ml", "grams" };
        RuleFor(f => f.Unit)
            .Must(f => allowedUnits.Contains(f))
            .WithMessage("unit must be either grams or ml");
    }
}
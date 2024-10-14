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
    private double _protein;

    public double Protein
    {
        get { return _protein; }
        set
        {
            _protein = value;
            CalculateKCal();
        }
    }

    private double _carbohydrate;

    public double Carbohydrate
    {
        get { return _carbohydrate; }
        set
        {
            _carbohydrate = value;
            CalculateKCal();
        }
    }

    private double _fat;

    public double Fat
    {
        get { return _fat; }
        set
        {
            _fat = value;
            CalculateKCal();
        }
    }

    private double _kCal;

    public double KCal
    {
        get { return _kCal; }
        set { _kCal = value; }
    }

    public string Unit { get; set; } = "grams";

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

    private void CalculateKCal()
    {
        _kCal = 4.0 * Protein + 4.0 * Carbohydrate + 9.0 * Fat;
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
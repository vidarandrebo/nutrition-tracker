using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems.Entities;

public class Micronutrient : BaseEntity
{
    public string Name { get; set; }
    public double Amount { get; set; }
    public MassUnit MassUnit { get; set; }

    public Micronutrient(string name, double amount, MassUnit massUnit)
    {
        Name = name;
        Amount = amount;
        MassUnit = massUnit;
    }
}
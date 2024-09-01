using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Domain.FoodItems.Entities;

/// <summary>
/// Macronutrients stores the amount of protein, carbohydrate and fat per 100grams of the item
/// </summary>
public class Macronutrients : BaseEntity
{
    public double Protein { get; set; }
    public double Carbohydrate { get; set; }
    public double Fat { get; set; }
    public double KCal { get; set; }


    public Macronutrients(double protein, double carbohydrate, double fat, double kCal)
    {
        Protein = protein;
        Carbohydrate = carbohydrate;
        Fat = fat;
        KCal = kCal;
    }

    public MacronutrientsDto ToDto()
    {
        return new MacronutrientsDto(Protein, Carbohydrate, Fat, KCal);
    }
}
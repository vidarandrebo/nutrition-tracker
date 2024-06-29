using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems;

/// <summary>
/// Macronutrients stores the amount of protein, carbohydrate and fat per 100grams of 
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
}
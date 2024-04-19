using Microsoft.EntityFrameworkCore;

namespace NutritionTracker.Domain.FoodItems;

[Owned]
public class NutritionalContent
{
    public double Protein { get; set; }
    public double Carbohydrate { get; set; }
    public double Fat { get; set; }
    public double KCal { get; set; }


    public NutritionalContent(double protein, double carbohydrate, double fat, double kCal)
    {
        Protein = protein;
        Carbohydrate = carbohydrate;
        Fat = fat;
        KCal = kCal;
    }
}
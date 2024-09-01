using System.Collections.Generic;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems.Dtos;

public class MacronutrientsDto(double protein, double carbohydrate, double fat, double kCal)
    : ValueObject
{
    public double Protein { get; set; } = protein;
    public double Carbohydrate { get; set; } = carbohydrate;
    public double Fat { get; set; } = fat;
    public double KCal { get; set; } = kCal;

    protected override IEnumerable<object> GetEqualityComponents()
    {
        yield return Protein;
        yield return Carbohydrate;
        yield return Fat;
        yield return KCal;
    }

    public static MacronutrientsDto operator +(MacronutrientsDto a, MacronutrientsDto b)
    {
        return new MacronutrientsDto(
            a.Protein + b.Protein,
            a.Carbohydrate + b.Carbohydrate,
            a.Fat + b.Fat,
            a.KCal + b.KCal
        );
    }

}
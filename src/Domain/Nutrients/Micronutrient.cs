using System.Collections.Generic;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Domain.Nutrients;

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

    public MicronutrientDto ToDto()
    {
        return new MicronutrientDto(Name, Amount, MassUnit);
    }
}

public static class MicronutrientIEnumerable
{
    public static List<MicronutrientDto> ToDto(this IEnumerable<Micronutrient> micronutrients)
    {
        var dtos = new List<MicronutrientDto>();
        foreach (var micronutrient in micronutrients)
        {
            dtos.Add(micronutrient.ToDto());
        }

        return dtos;
    }
}
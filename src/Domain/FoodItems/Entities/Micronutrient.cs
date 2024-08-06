using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;
using System.Collections.Generic;

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
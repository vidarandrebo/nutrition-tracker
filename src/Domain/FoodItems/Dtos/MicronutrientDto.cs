using System.Collections.Generic;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems.Dtos;

public class MicronutrientDto : ValueObject
{
    public string Name { get; set; }
    public double Amount { get; set; }
    public MassUnit MassUnit { get; set; }

    protected override IEnumerable<object> GetEqualityComponents()
    {
        yield return Name;
        yield return Amount;
        yield return MassUnit;
    }
}

public static class MicronutrientDtoIEnumerable
{
    public static void Merge(this IEnumerable<MicronutrientDto> a, IEnumerable<MicronutrientDto> b)
    {
        
    }
}
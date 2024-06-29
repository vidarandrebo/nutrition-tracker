using System;

namespace NutritionTracker.Domain.FoodItems
{
    public record FoodItemDTO(
        Guid Id,
        string Brand,
        string ProductName,
        Macronutrients Macronutrients,
        Guid OwnerId);
}
using System;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Recipes;

public class Ingredient : BaseEntity
{
    public double Amount { get; set; }
    public Guid FoodItemId { get; set; }
    public IngredientType Type { get; set; }
}
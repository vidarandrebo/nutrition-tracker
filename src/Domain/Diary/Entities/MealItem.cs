using System;
using System.Collections.Generic;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.Nutrients;

namespace NutritionTracker.Domain.Diary.Entities;

public class MealItem : BaseEntity
{
    public Guid RecipeId { get; set; }
    public Guid FoodItemId { get; set; }
    public Macronutrients Macronutrients { get; set; }
    public List<Micronutrient> Micronutrients { get; set; }
}

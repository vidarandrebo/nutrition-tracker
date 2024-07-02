using System;
using System.Collections.Generic;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems;

namespace NutritionTracker.Domain.Accounts;

public class Account : BaseEntity
{
    public Macronutrients MacronutrientGoal { get; set; }
    public List<Micronutrient> MicronutrientGoals { get; set; }

    public Account()
    {
    }
}
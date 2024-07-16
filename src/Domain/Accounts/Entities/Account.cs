using System;
using System.Collections.Generic;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Entities;

namespace NutritionTracker.Domain.Accounts.Entities;

public class Account : BaseEntity
{
    public Macronutrients MacronutrientGoal { get; set; }
    public List<Micronutrient> MicronutrientGoals { get; set; }

    public Account(Guid id)
    {
        Id = id;
        MacronutrientGoal = new Macronutrients(0.0, 0.0, 0.0, 0.0);
        MicronutrientGoals = new List<Micronutrient>();
    }
}
using NutritionTracker.Domain.Accounts.Dtos;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Entities;
using System;
using System.Collections.Generic;

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

    public AccountDto ToDto()
    {
        var dto = new AccountDto(
            Id,
            MacronutrientGoal.ToDto(),
            MicronutrientGoals.ToDto());
        return dto;
    }
}
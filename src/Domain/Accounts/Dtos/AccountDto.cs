using NutritionTracker.Domain.FoodItems.Dtos;
using System;
using System.Collections.Generic;

namespace NutritionTracker.Domain.Accounts.Dtos
{
    public class AccountDto
    {
        public Guid Id { get; set; }
        public MacronutrientsDto MacronutrientGoal { get; set; }
        public List<MicronutrientDto> MicronutrientGoals { get; set; }
        public AccountDto(Guid id, MacronutrientsDto macronutrientGoal, List<MicronutrientDto> micronutrientGoals)
        {
            Id = id;
            MacronutrientGoal = macronutrientGoal;
            MicronutrientGoals = micronutrientGoals;
        }
    }
}

using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;
using System;
using System.Collections.Generic;

namespace NutritionTracker.Domain.Recipes.Entities;

public class Recipe : BaseEntity
{
    public List<Ingredient> Ingredients { get; set; }

    public Recipe()
    {
    }

    public List<MicronutrientDto> Micronutrients()
    {
        return new List<MicronutrientDto>();
    }

    public MacronutrientsDto Macronutrients()
    {
        var macronutrients = new MacronutrientsDto(0.0, 0.0, 0.0, 0.0);
        foreach (var ingredient in Ingredients)
        {
            //macroNutrients += ingredient.Macronutrients();
        }

        return macronutrients;
    }

    public void ResolveDependencyCycles()
    {
        throw new NotImplementedException();
    }
}
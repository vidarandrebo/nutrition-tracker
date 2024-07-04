using System;
using System.Collections.Generic;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Domain.Recipes.Entities;

public class Recipe : BaseEntity
{
    public List<Ingredient> Ingredients { get; set; }
    public Recipe()
    {
        
    }

    public MicronutrientDto Micronutrients()
    {
        return new MicronutrientDto();
    }

    public MacronutrientsDto Macronutrients()
    {
        var macroNutrients = new MacronutrientsDto(0.0,0.0,0.0,0.0);
        foreach (var ingredient in Ingredients)
        {
            //macroNutrients += ingredient.Macronutrients();
        }

        return macroNutrients;
    }

    public void ResolveDepenencyCycles()
    {
        throw new NotImplementedException();
    }
}
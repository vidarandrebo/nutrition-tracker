using System.Collections.Generic;
using FluentResults;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;
using NutritionTracker.Domain.FoodItems.Entities;
using NutritionTracker.Domain.Recipes.Exceptions;

namespace NutritionTracker.Domain.Recipes.Entities;

public class Ingredient : BaseEntity
{
    public double Amount { get; set; }
    private FoodItem? _foodItem;
    private Recipe? _recipe;

    /// <summary>
    /// The FoodItem for the ingredient
    /// </summary>
    /// <exception cref="MultipleIngredientSourceException">Throws exception if both FoodItem and Recipe have a non-null value</exception>
    public FoodItem? FoodItem
    {
        get { return _foodItem; }
        set
        {
            if (Recipe is null)
            {
                _foodItem = value;
            }
            else
            {
                throw new MultipleIngredientSourceException(
                    "An ingredient cannot have both a recipe and and foodItem as a source");
            }
        }
    }

    /// <summary>
    /// The Recipe for the ingredient
    /// </summary>
    /// <exception cref="MultipleIngredientSourceException">Throws exception if both FoodItem and Recipe have a non-null value</exception>
    public Recipe? Recipe
    {
        get { return _recipe; }
        set
        {
            if (FoodItem is null)
            {
                _recipe = value;
            }
            else
            {
                throw new MultipleIngredientSourceException(
                    "An ingredient cannot have both a recipe and and foodItem as a source");
            }
        }
    }

    public List<MicronutrientDto> Micronutrients()
    {
        return new List<MicronutrientDto>();
    }

    public Result<MacronutrientsDto> Macronutrients()
    {
        if (FoodItem is not null)
        {
            return Result.Ok(FoodItem.Macronutrients.ToDto());
        }

        if (Recipe is not null)
        {
            return Result.Ok(Recipe.Macronutrients());
        }

        return Result.Fail("Ingredient has no foodItem or recipe");
    }
}
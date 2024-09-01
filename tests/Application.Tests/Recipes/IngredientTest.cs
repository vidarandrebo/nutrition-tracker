using NutritionTracker.Domain.FoodItems.Entities;
using NutritionTracker.Domain.Recipes.Entities;
using NutritionTracker.Domain.Recipes.Exceptions;

namespace NutritionTracker.Application.Tests.Recipes;

public class IngredientTest
{
    [Fact]
    public void ThrowIfFoodItemAndRecipe()
    {
        var ingredient = new Ingredient();
        var foodItem = new FoodItem(
            "food brand",
            "prod name",
            new Macronutrients(1.0, 1.0, 1.0, 1.0),
            Guid.Empty,
            new List<Micronutrient>());

        var recipe = new Recipe();
        ingredient.FoodItem = foodItem;
        Assert.Throws<MultipleIngredientSourceException>(() => { ingredient.Recipe = recipe; });
    }

    [Fact]
    public void ThrowIfRecipeAndFoodItem()
    {
        var ingredient = new Ingredient();
        var foodItem = new FoodItem(
            "food brand",
            "prod name",
            new Macronutrients(1.0, 1.0, 1.0, 1.0),
            Guid.Empty,
            new List<Micronutrient>());

        var recipe = new Recipe();
        ingredient.Recipe = recipe;
        Assert.Throws<MultipleIngredientSourceException>(() => { ingredient.FoodItem = foodItem; });
    }
}
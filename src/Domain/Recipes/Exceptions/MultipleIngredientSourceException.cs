using System;

namespace NutritionTracker.Domain.Recipes.Exceptions;

public class MultipleIngredientSourceException : Exception
{
    public MultipleIngredientSourceException()
    {
    }

    public MultipleIngredientSourceException(string message) : base(message)
    {
    }

    public MultipleIngredientSourceException(string message, Exception inner) : base(message, inner)
    {
    }
}
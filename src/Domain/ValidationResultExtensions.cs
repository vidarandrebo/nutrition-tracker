using System.Collections.Generic;
using FluentResults;
using FluentValidation.Results;

namespace NutritionTracker.Domain;

public static class ValidationResultExtensions
{
    public static List<Error> GetErrorList(this ValidationResult validationResult)
    {
        var errors = new List<Error>();
        foreach (var err in validationResult.Errors)
        {
            errors.Add(new Error(err.ErrorMessage));
        }

        return errors;
    }
}
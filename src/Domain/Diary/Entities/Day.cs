using System;
using System.Collections.Generic;
using System.Linq;
using FluentResults;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Diary.Entities;

public class Day : BaseEntity
{
    public DateOnly Date { get; set; }
    public double ActivityCalories { get; set; }
    public List<Meal> Meals { get; set; }

    public Day()
    {
        Id = Guid.NewGuid();
        Date = DateOnly.FromDateTime(DateTime.Now);
        Meals = new List<Meal>();
    }

    public void AddMeal(Meal meal)
    {
        if (Meals.Count > 0)
        {
            meal.SequenceNumber = Meals.Max(x => x.SequenceNumber) + 1;
        }
        else
        {
            meal.SequenceNumber = 0;
        }

        Meals.Add(meal);
    }
}
using System;
using System.Collections.Generic;
using System.Linq;
using FluentResults;

namespace NutritionTracker.Domain.Diary.Entities;

public class Day
{
    public DateOnly Date { get; set; }
    public double ActivityCalories { get; set; }
    public List<Meal> Meals { get; set; }

    public Day()
    {
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
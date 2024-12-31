using System.Collections.Generic;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Diary.Entities;

public class Meal : BaseEntity
{
    public long SequenceNumber { get; set; }
    public List<MealItem> MealItems { get; set; }

    public Meal()
    {
        MealItems = new List<MealItem>();
    }

    public string MealTitle
    {
        get { return $"Meal {SequenceNumber + 1}"; }
    }
}
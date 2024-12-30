using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.Diary.Entities;

public class Meal : BaseEntity
{
    public long SequenceNumber { get; set; }

    public Meal()
    {
    }

    public string MealTitle
    {
        get
        {
            return $"Meal {SequenceNumber + 1}";
        }
    }
}
using System;
using Microsoft.AspNetCore.Components;

namespace NutritionTracker.Web.Pages.Diary.Meals;

public partial class Index : ComponentBase
{
    [Parameter] public string MealId { get; set; }
}
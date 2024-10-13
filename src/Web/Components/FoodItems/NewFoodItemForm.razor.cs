using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;
using NutritionTracker.Domain.FoodItems.Contracts;

namespace NutritionTracker.Web.Components.FoodItems;

public partial class NewFoodItemForm : ComponentBase
{
    [Inject] private ILogger<NewFoodItemForm> Logger { get; set; }
    [Parameter] public PostFoodItemRequest Model { get; set; }
    [Parameter] public EventCallback OnValidSubmit { get; set; }
}
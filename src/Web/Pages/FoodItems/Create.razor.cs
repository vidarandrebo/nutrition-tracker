using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Domain.FoodItems.Contracts;

namespace NutritionTracker.Web.Pages.FoodItems;

public partial class Create
{
    [SupplyParameterFromForm] private PostFoodItemRequest FoodItemRequest { get; set; } = new();

    private async Task CreateFoodItem()
    {
        var user = await UserAccessor.GetRequiredUserAsync(SignInManager.Context);
        var response = await Mediator.Send(new AddFoodItem.Request(FoodItemRequest, user.AccountId));
        if (response.IsSuccess)
        {
            Logger.LogInformation("Created fooditem with brand {brand}", FoodItemRequest.Brand);
            NavigationManager.NavigateTo("/fooditems");
        }
        else
        {
            Logger.LogError(response.Errors.FirstOrDefault().ToString());
        }
    }
}
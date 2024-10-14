using System.Linq;
using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Domain.FoodItems.Contracts;
using NutritionTracker.Web.Identity;

namespace NutritionTracker.Web.Pages.FoodItems;

public partial class Create
{
    [CascadingParameter] private HttpContext? HttpContext { get; set; }
    [Inject] private ILogger<Create> Logger { get; set; }
    public PostFoodItemRequest FoodItemRequest { get; set; } = new();
    [Inject] private NavigationManager NavigationManager { get; set; } = null!;
    [Inject] private IdentityUserAccessor UserAccessor { get; set; } = null!;
    [Inject] private IMediator Mediator { get; set; } = null!;
    [Inject] private AuthenticationStateProvider AuthenticationStateProvider { get; set; } = null!;

    private async Task CreateFoodItem()
    {
        Logger.LogInformation("Creating new food item");
        var userResponse = await UserAccessor.GetUserFromAuthStateAsync();
        if (userResponse.IsFailed)
        {
            Logger.LogError("Failed to get user response");
            return;
        }
        var user  = userResponse.Value;
        Logger.LogInformation(user.AccountId.ToString());
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
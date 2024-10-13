using System;
using System.Threading.Tasks;
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Web.Pages.FoodItems;

public partial class Index
{
    private string? _searchTerm;

    public string? SearchTerm
    {
        get { return _searchTerm; }
        set { _searchTerm = value; }
    }

    public FoodItemDto[] Items { get; set; } = Array.Empty<FoodItemDto>();

    protected override async Task OnInitializedAsync()
    {
        var getFoodItemsResult = await Mediator.Send(new GetFoodItems.Request());
        if (getFoodItemsResult.IsSuccess)
        {
            Items = getFoodItemsResult.Value;
        }

        await base.OnInitializedAsync();
    }
}
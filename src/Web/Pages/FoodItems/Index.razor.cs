using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Components;
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Application.FoodItems.Commands;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Web.Pages.FoodItems;

public partial class Index
{
    [Inject] private IMediator Mediator { get; set; }
    private string? _searchTerm;

    public string? SearchTerm
    {
        get { return _searchTerm; }
        set
        {
            _searchTerm = value;
            if (SearchTerm != null)
            {
                FilteredItems = _items.Where(i => i.ProductName.Contains(SearchTerm));
            }
        }
    }


    private IEnumerable<FoodItemDto>? _filteredItems;

    public IEnumerable<FoodItemDto>? FilteredItems
    {
        get
        {
            if (_filteredItems is not null)
            {
                return _filteredItems.Take(100);
            }

            return null;
        }
        set
        {
            _filteredItems = value;
        }
    }


    private FoodItemDto[] _items { get; set; } = Array.Empty<FoodItemDto>();

    protected override async Task OnInitializedAsync()
    {
        var getFoodItemsResult = await Mediator.Send(new GetFoodItems.Request());
        if (getFoodItemsResult.IsSuccess)
        {
            _items = getFoodItemsResult.Value;
        }

        FilteredItems = _items;

        await base.OnInitializedAsync();
    }
}
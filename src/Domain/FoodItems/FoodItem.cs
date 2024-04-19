using System;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems;

public class FoodItem : BaseEntity
{
    public string Brand { get; set; }
    public string ProductName { get; set; }
    public NutritionalContent NutritionalContent { get; set; }
    public Guid OwnerId { get; set; }


    // EF Core
    private FoodItem()
    {
        Brand = "";
        ProductName = "";
        NutritionalContent = new NutritionalContent(0.0, 0.0, 0.0, 0.0);
    }

    public FoodItem(string brand, string productName, NutritionalContent nutritionalContent, Guid ownerId)
    {
        Id = Guid.NewGuid();
        Brand = brand;
        ProductName = productName;
        NutritionalContent = nutritionalContent;
        OwnerId = ownerId;
    }

    public FoodItemDTO ToDTO()
    {
        return new FoodItemDTO(Id, Brand, ProductName, NutritionalContent, OwnerId);
    }
}
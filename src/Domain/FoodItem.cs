using System;
using Domain.Common;

namespace Domain;

public class FoodItem : BaseEntity
{
    public string Brand { get; set; }
    public string ProductName { get; set; }
    public NutritionalContent NutritionalContent { get; set; }
    public Guid OwnerId { get; set; }


    // EF Core
    private FoodItem()
    {
    }

    public FoodItem(string brand, string productName, NutritionalContent nutritionalContent, Guid ownerId)
    {
        Id = Guid.NewGuid();
        Brand = brand;
        ProductName = productName;
        NutritionalContent = nutritionalContent;
        OwnerId = ownerId;
    }
}
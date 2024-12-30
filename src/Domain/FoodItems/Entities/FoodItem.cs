using System;
using System.Collections.Generic;
using System.Linq;
using NutritionTracker.Domain.Common;
using NutritionTracker.Domain.FoodItems.Dtos;

namespace NutritionTracker.Domain.FoodItems.Entities;

public class FoodItem : BaseEntity
{
    public string Brand { get; set; }
    public string ProductName { get; set; }
    public Macronutrients Macronutrients { get; set; }
    public List<Micronutrient> Micronutrients { get; set; }
    public Guid OwnerId { get; set; }


    // EF Core
    private FoodItem()
    {
        Brand = "";
        ProductName = "";
        Macronutrients = new Macronutrients(0.0, 0.0, 0.0, 0.0);
        Micronutrients = new List<Micronutrient>();
    }

    public FoodItem(string brand, string productName, Macronutrients macronutrients, Guid ownerId,
        IEnumerable<Micronutrient> micronutrients)
    {
        Id = Guid.NewGuid();
        Brand = brand;
        ProductName = productName;
        Macronutrients = macronutrients;
        OwnerId = ownerId;
        Micronutrients = micronutrients.ToList();
    }

    public FoodItemDto ToDTO()
    {
        return new FoodItemDto(Id, Brand, ProductName, Macronutrients.ToDto(), OwnerId);
    }
}
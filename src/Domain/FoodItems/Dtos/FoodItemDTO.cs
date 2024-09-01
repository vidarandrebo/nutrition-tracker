using System;
using System.Collections.Generic;
using NutritionTracker.Domain.Common;

namespace NutritionTracker.Domain.FoodItems.Dtos
{
    public class FoodItemDto : ValueObject
    {
        public Guid Id { get; set; }
        public string Brand { get; set; }
        public string ProductName { get; set; }
        public MacronutrientsDto Macronutrients { get; set; }
        public Guid Owner { get; set; }

        public FoodItemDto(Guid id, string brand, string productName, MacronutrientsDto macronutrients, Guid owner)
        {
            Id = id;
            Brand = brand;
            ProductName = productName;
            Macronutrients = macronutrients;
            Owner = owner;
        }

        protected override IEnumerable<object> GetEqualityComponents()
        {
            yield return Id;
            yield return Brand;
            yield return ProductName;
            yield return Macronutrients;
            yield return Owner;
        }
    }
}
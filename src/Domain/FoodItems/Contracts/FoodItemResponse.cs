using NutritionTracker.Domain.FoodItems.Dtos;
using System;
using System.Collections.Generic;

namespace NutritionTracker.Domain.FoodItems.Contracts
{
    public class FoodItemResponse
    {
        public Guid Id { get; set; }
        public string Brand { get; set; }
        public string ProductName { get; set; }
        public MacronutrientsResponse Macronutrients { get; set; }
        public Guid Owner { get; set; }
        public class MacronutrientsResponse
        {
            public double Protein { get; set; }
            public double Carbohydrate { get; set; }
            public double Fat { get; set; }
            public double KCal { get; set; }
            public MacronutrientsResponse(double protein, double carbohydrate, double fat, double kCal)
            {
                Protein = protein;
                Carbohydrate = carbohydrate;
                Fat = fat;
                KCal = kCal;
            }
        }
        public FoodItemResponse(Guid id, string brand, string productName, MacronutrientsResponse macronutrients, Guid owner)
        {
            Id = id;
            Brand = brand;
            ProductName = productName;
            Macronutrients = macronutrients;
            Owner = owner;
        }
        public static FoodItemResponse[] FromDtos(FoodItemDto[] dtos)
        {
            var responses = new List<FoodItemResponse>();
            foreach (var dto in dtos)
            {
                responses.Add(FromDto(dto));
            }
            return responses.ToArray();
        }
        public static FoodItemResponse FromDto(FoodItemDto dto)
        {
            var response = new FoodItemResponse(
                dto.Id,
                dto.Brand,
                dto.ProductName,
                new MacronutrientsResponse(
                    dto.Macronutrients.Protein,
                    dto.Macronutrients.Carbohydrate,
                    dto.Macronutrients.Fat,
                    dto.Macronutrients.KCal
                    ),
                dto.Owner
                );
            return response;
        }
    }
}

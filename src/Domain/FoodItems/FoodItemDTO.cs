using System;

namespace Domain.FoodItems
{
    public record FoodItemDTO(Guid Id, string Brand, string ProductName, NutritionalContent NutritionalContent, Guid OwnerId);
}

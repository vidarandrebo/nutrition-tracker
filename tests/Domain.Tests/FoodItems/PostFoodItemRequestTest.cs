using System.Linq;
using NutritionTracker.Domain.FoodItems.Contracts;

namespace NutritionTracker.Domain.Tests.FoodItems;

public class PostFoodItemRequestTest
{
    public FoodItemValidator FoodItemFormFoodItemValidator;

    public PostFoodItemRequestTest()
    {
        FoodItemFormFoodItemValidator = new FoodItemValidator();
    }

    public PostFoodItemRequest DefaultForm()
    {
        var form = new PostFoodItemRequest();
        form.Brand = "TestBrand";
        form.ProductName = "ProductName";
        form.Protein = 20;
        form.Carbohydrate = 20;
        form.Fat = 20;
        form.KCal = 340;
        form.Unit = "grams";
        return form;
    }

    [Fact]
    public void CorrectForm()
    {
        var form = DefaultForm();
        var validationResult = FoodItemFormFoodItemValidator.Validate(form);
        Assert.True(validationResult.IsValid);
    }

    [Fact]
    public void IncorrectUnit()
    {
        var form = DefaultForm();
        form.Unit = "dl";
        var validationResult = FoodItemFormFoodItemValidator.Validate(form);
        var errorMsgs = validationResult.Errors.Select(e => e.ErrorMessage).ToList();
        Assert.Contains("unit must be either grams or ml", errorMsgs);
    }
}
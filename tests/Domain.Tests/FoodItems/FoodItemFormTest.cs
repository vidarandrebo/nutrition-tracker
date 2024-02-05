using System.Linq;
using Domain.FoodItems;

namespace Domain.Tests.FoodItems;

public class FoodItemFormTest
{
    public FoodItemValidator FoodItemFormFoodItemValidator;

    public FoodItemFormTest()
    {
        FoodItemFormFoodItemValidator = new FoodItemValidator();
    }

    public FoodItemForm DefaultForm()
    {
        var form = new FoodItemForm();
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
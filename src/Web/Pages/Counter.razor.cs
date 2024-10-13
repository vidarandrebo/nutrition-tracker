using Microsoft.Extensions.Logging;

namespace NutritionTracker.Web.Pages;

public partial class Counter
{
    private int currentCount = 0;
    private string? _textValue;

    public string? TextValue
    {
        get { return _textValue; }
        set
        {
            _textValue = value;
            Logger.LogInformation(TextValue);
        }
    }

    private void IncrementCount()
    {
        currentCount++;
    }
}
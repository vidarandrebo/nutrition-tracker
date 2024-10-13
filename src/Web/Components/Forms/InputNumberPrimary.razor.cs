using System.Collections.Generic;
using Microsoft.AspNetCore.Components;

namespace NutritionTracker.Web.Components.Forms;

public partial class InputNumberPrimary<TValue>
{
    [Parameter] public TValue? Value { get; set; }
    [Parameter(CaptureUnmatchedValues = true)]
    public Dictionary<string, object>? InputAttributes { get; set; }
}
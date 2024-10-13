using System.Threading.Tasks;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Http;

namespace NutritionTracker.Web.Pages.Account.Manage;

public partial class PersonalData
{
    [CascadingParameter] private HttpContext HttpContext { get; set; } = default!;

    protected override async Task OnInitializedAsync()
    {
        _ = await UserAccessor.GetRequiredUserAsync(HttpContext);
    }
}
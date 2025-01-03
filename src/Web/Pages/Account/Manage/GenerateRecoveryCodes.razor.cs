using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Logging;
using NutritionTracker.Infrastructure.Identity;

namespace NutritionTracker.Web.Pages.Account.Manage;

public partial class GenerateRecoveryCodes
{
    private string? message;
    private ApplicationUser user = default!;
    private IEnumerable<string>? recoveryCodes;

    [CascadingParameter] private HttpContext HttpContext { get; set; } = default!;

    protected override async Task OnInitializedAsync()
    {
        user = await UserAccessor.GetRequiredUserAsync(HttpContext);

        var isTwoFactorEnabled = await UserManager.GetTwoFactorEnabledAsync(user);
        if (!isTwoFactorEnabled)
        {
            throw new InvalidOperationException(
                "Cannot generate recovery codes for user because they do not have 2FA enabled.");
        }
    }

    private async Task OnSubmitAsync()
    {
        var userId = await UserManager.GetUserIdAsync(user);
        recoveryCodes = await UserManager.GenerateNewTwoFactorRecoveryCodesAsync(user, 10);
        message = "You have generated new recovery codes.";

        Logger.LogInformation("User with ID '{UserId}' has generated new 2FA recovery codes.", userId);
    }
}
using System.Threading.Tasks;
using FluentResults;
using Microsoft.AspNetCore.Components.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using NutritionTracker.Infrastructure.Identity;

namespace NutritionTracker.Web.Identity;

internal sealed class IdentityUserAccessor
{
    private readonly UserManager<ApplicationUser> _userManager;
    private readonly IdentityRedirectManager _redirectManager;
    private readonly AuthenticationStateProvider _authenticationStateProvider;

    public IdentityUserAccessor(UserManager<ApplicationUser> userManager,
        IdentityRedirectManager redirectManager, AuthenticationStateProvider authenticationStateProvider)
    {
        _userManager = userManager;
        _redirectManager = redirectManager;
        _authenticationStateProvider = authenticationStateProvider;
    }

    public async Task<ApplicationUser> GetRequiredUserAsync(HttpContext context)
    {
        var user = await _userManager.GetUserAsync(context.User);

        if (user is null)
        {
            _redirectManager.RedirectToWithStatus("Account/InvalidUser",
                $"Error: Unable to load user with ID '{_userManager.GetUserId(context.User)}'.", context);
        }

        return user;
    }

    public async Task<Result<ApplicationUser>> GetUserFromAuthStateAsync()
    {
        var userState = await _authenticationStateProvider.GetAuthenticationStateAsync();
        var user = await _userManager.GetUserAsync(userState.User);
        if (user is null)
        {
            return Result.Fail("No user found.");
        }

        return Result.Ok(user);
    }
}
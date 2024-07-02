using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using FluentResults;
using MediatR;
using Microsoft.AspNetCore.Identity;
using NutritionTracker.Application.Identity;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Infrastructure.Identity;

public class IdentityService : IIdentityService
{
    private readonly UserManager<ApplicationUser> _userManager;
    private readonly SignInManager<ApplicationUser> _signInManager;

    public IdentityService(UserManager<ApplicationUser> userManager, SignInManager<ApplicationUser> signInManager)
    {
        _userManager = userManager;
        _signInManager = signInManager;
    }

    public async Task<Result<ApplicationUserDto>> LoginUser(string email, string password)
    {
        var user = await _userManager.FindByEmailAsync(email);
        if (user is null)
        {
            return Result.Fail((new Error("User not found")));
        }

        var correctPasswd = await _userManager.CheckPasswordAsync(user, password);
        if (!correctPasswd)
        {
            return Result.Fail(new Error("Password is wrong"));
        }

        if (user.UserName is null)
        {
            return Result.Fail(new Error("Username is null"));
        }

        return Result.Ok(new ApplicationUserDto());
    }

    public async Task<Unit> LogoutUser()
    {
        await _signInManager.SignOutAsync();
        return Unit.Value;
    }

    public async Task<Result<Guid>> RegisterUser(string email, string password)
    {
        var user = new ApplicationUser();
        user.Email = email;
        
        var registerResult = await _userManager.CreateAsync(user, password);
        var errors = new List<Error>();
        foreach (var err in registerResult.Errors)
        {
            errors.Add(new Error(err.Description));
        }

        if (errors.Count > 0)
        {
            return Result.Fail(errors);
        }

        return Result.Ok(user.Id);
    }
}
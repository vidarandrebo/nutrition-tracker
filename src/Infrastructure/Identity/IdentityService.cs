using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;
using FluentResults;
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Infrastructure.Identity;

public class IdentityService : IIdentityService
{
    private readonly UserManager<ApplicationUser> _userManager;
    private readonly SignInManager<ApplicationUser> _signInManager;
    private readonly IConfiguration _configuration;
    private readonly ITokenHandler _tokenHandler;

    public IdentityService(UserManager<ApplicationUser> userManager, SignInManager<ApplicationUser> signInManager, IConfiguration configuration, ITokenHandler tokenHandler)
    {
        _userManager = userManager;
        _signInManager = signInManager;
        _configuration = configuration;
        _tokenHandler = tokenHandler;
    }

    public async Task<Result<AccessTokenResponse>> LoginUser(string email, string password)
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

        // This should not happen as user is resolved by email...
        if (user.Email is null)
        {
            return Result.Fail("Email is null");
        }

        var accessToken = _tokenHandler.AccessToken(user.Id, user.Email);
        var refreshToken = _tokenHandler.RefreshToken(user.Id, user.Email);

        return Result.Ok(new AccessTokenResponse()
        {
            AccessToken = accessToken,
            RefreshToken = refreshToken,
            ExpiresIn = 60*60*24,
        });
        
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
        user.UserName = email;
        
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

        return Result.Ok(user.AccountId);
    }

    public Result<Guid> GetUserIdFromRequest(HttpContext context)
    {
        var claimsResponse = GetJwtClaimsFromRequest(context);
        if (claimsResponse.IsFailed)
        {
            return Result.Fail("Token not provided");
        }
        var id = claimsResponse.Value.Claims.FirstOrDefault(c => c.Type == ClaimTypes.NameIdentifier)?.Value ?? "";
        var parsedGuid = GuidHelper.GuidOrEmpty(id);
        if (parsedGuid == Guid.Empty)
        {
            return Result.Fail(new Error("Id not found in claims"));
        }

        return Result.Ok(parsedGuid);
    }
    private Result<ClaimsPrincipal> GetJwtClaimsFromRequest(HttpContext context)
    {
        var token = context.Request.Headers["Authorization"].FirstOrDefault()?.Split(" ").Last();
        if (token is null)
        {
            return Result.Fail("Token not provided");
        }

        var validatedTokenResult = _tokenHandler.ValidateToken(token);
        if (validatedTokenResult.IsSuccess)
        {
            return Result.Ok(validatedTokenResult.Value);
        }

        return Result.Fail(validatedTokenResult.Errors.ToString());
    }
    public  Result<string> GetUserNameFromRequest(HttpContext context)
    {
        var claimsResult = GetJwtClaimsFromRequest(context);
        if (claimsResult.IsSuccess)
        {
            var username =  claimsResult.Value.Claims.FirstOrDefault(c => c.Type == ClaimTypes.Name)?.Value;
            if (username is not null)
            {
                return Result.Ok(username);
            }

            return Result.Fail("username not found in claims");
        }

        return Result.Fail(claimsResult.Errors.ToString());
    }
}
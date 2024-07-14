using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;
using FluentResults;
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
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

    public IdentityService(UserManager<ApplicationUser> userManager, SignInManager<ApplicationUser> signInManager, IConfiguration configuration)
    {
        _userManager = userManager;
        _signInManager = signInManager;
        _configuration = configuration;
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

        var tokenHandler = new JwtSecurityTokenHandler();

        var accessToken = AccessToken(user.Id, user.Email);
        var refreshToken = RefreshToken(user.Id, user.Email);
        


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

    public string RefreshToken(Guid id, string email)
    {
        var tokenHandler = new JwtSecurityTokenHandler();
        var token = CreateToken(id, email, DateTime.Now.AddDays(30));
        return tokenHandler.WriteToken(token);
    }
    public string AccessToken(Guid id, string email)
    {
        var tokenHandler = new JwtSecurityTokenHandler();
        var token = CreateToken(id, email, DateTime.Now.AddDays(1));
        return tokenHandler.WriteToken(token);
    }
    private JwtSecurityToken CreateToken(Guid id, string email, DateTime expires)
    {
        var claimList = new List<Claim>();
        claimList.Add(new Claim(ClaimTypes.Email, email));
        claimList.Add(new Claim(ClaimTypes.NameIdentifier, id.ToString()));
        var signKey = _configuration["Jwt:Secret"];
        if (signKey is null)
        {
            throw new Exception("Jwt secret key not set");
        }
        var token = new JwtSecurityToken(
            claims: claimList,
            expires: DateTime.Now.AddDays(30),
            signingCredentials: new SigningCredentials(new SymmetricSecurityKey(Encoding.UTF8.GetBytes(signKey)),
                SecurityAlgorithms.HmacSha256)
            , audience: _configuration["Jwt:ValidAudience"], issuer:_configuration["Jwt:ValidIssuer"]
        );
        return token;
    }
}
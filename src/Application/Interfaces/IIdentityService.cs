using System;
using System.IdentityModel.Tokens.Jwt;
using System.Threading.Tasks;
using FluentResults;
using Microsoft.AspNetCore.Authentication.BearerToken;
using NutritionTracker.Application.Identity;

namespace NutritionTracker.Application.Interfaces;

public interface IIdentityService
{

    public string AccessToken(Guid id, string email);
    public string RefreshToken(Guid id, string email);
    public Task<Result<Guid>> RegisterUser(string email, string password);
    public Task<Result<AccessTokenResponse>> LoginUser(string email, string password);
}
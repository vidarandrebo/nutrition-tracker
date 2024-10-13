using System;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using FluentResults;
using Microsoft.AspNetCore.Http;

namespace NutritionTracker.Application.Interfaces;

public interface ITokenHandler
{
    string AccessToken(Guid id, string email);
    string RefreshToken(Guid id, string email);
    Result<ClaimsPrincipal> ValidateToken(string token);
}
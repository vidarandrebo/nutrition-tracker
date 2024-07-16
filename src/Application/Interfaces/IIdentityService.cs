using System;
using System.Threading.Tasks;
using FluentResults;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Http;

namespace NutritionTracker.Application.Interfaces;

public interface IIdentityService
{

    Task<Result<Guid>> RegisterUser(string email, string password);
    Task<Result<AccessTokenResponse>> LoginUser(string email, string password);
    Result<Guid> GetUserIdFromRequest(HttpContext context);
    Result<string> GetUserNameFromRequest(HttpContext context);
}
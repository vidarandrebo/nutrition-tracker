using System;
using FluentResults;
using Microsoft.AspNetCore.Http;

namespace NutritionTracker.Application.Interfaces;

public interface ITokenHandler
{
    
    string CreateToken(Guid id, string userName);
    Result<Guid> GetUserIdFromRequest(HttpContext context);
    string? GetUserNameFromRequest(HttpContext context);
}
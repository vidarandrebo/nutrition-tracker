using System;
using System.Security.Claims;
using FluentResults;
using Microsoft.AspNetCore.Http;

namespace Server;

public static class HttpContextExtensions
{
    public static Result<Guid> GetUserId(this HttpContext ctx)
    {
        var userId = ctx.User.FindFirstValue(ClaimTypes.NameIdentifier);
        var result = (userId != null) ? Result.Ok(Guid.Parse(userId)) : Result.Fail(new Error("User not logged in"));
        return result;
    }

    public static Result<string> GetUserName(this HttpContext ctx)
    {
        var userName = ctx.User.Identity?.Name;
        var result = (userName != null) ? Result.Ok(userName) : Result.Fail(new Error("User not logged in"));
        return result;
    }
}
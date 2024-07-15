using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Security.Claims;
using System.Text;
using FluentResults;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Infrastructure.Identity;

public class TokenHandler : ITokenHandler
{
    private readonly IConfiguration _configuration;

    public TokenHandler(IConfiguration configuration)
    {
        _configuration = configuration;
    }

    public string? GetUserNameFromRequest(HttpContext context)
    {
        var token = _getTokenFromRequest(context);
        return token?.Claims.FirstOrDefault(c => c.Type == ClaimTypes.Name)?.Value;
    }

    public Result<Guid> GetUserIdFromRequest(HttpContext context)
    {
        var token = _getTokenFromRequest(context);
        var id = token?.Claims.FirstOrDefault(c => c.Type == ClaimTypes.NameIdentifier)?.Value ?? "";
        var parsedGuid = GuidHelper.GuidOrEmpty(id);
        if (parsedGuid == Guid.Empty)
        {
            return Result.Fail(new Error("Failed to find ID in http-context"));
        }

        return Result.Ok(parsedGuid);
    }

    public string CreateToken(Guid id, string userName)
    {
        var jwtHandler = new JwtSecurityTokenHandler();
        var token = _createToken(id, userName);
        return jwtHandler.WriteToken(token);
    }

    private JwtSecurityToken? _getTokenFromRequest(HttpContext context)
    {
        var token = context.Request.Headers["Authorization"].FirstOrDefault()?.Split(" ").Last();
        if (token is null)
        {
            return null;
        }

        var validatedToken = ValidateToken(token);
        if (validatedToken is not null)
        {
            return validatedToken;
        }

        return null;
    }

    private JwtSecurityToken _createToken(Guid id, string userName)
    {
        var claimList = new List<Claim>();
        claimList.Add(new Claim(ClaimTypes.Name, userName));
        claimList.Add(new Claim(ClaimTypes.NameIdentifier, id.ToString()));
        var signKey = _configuration.GetValue<string>("Jwt:Secret") ?? throw new ArgumentNullException();
        var token = new JwtSecurityToken(
            claims: claimList,
            expires: DateTime.Now.AddDays(30),
            signingCredentials: new SigningCredentials(new SymmetricSecurityKey(Encoding.UTF8.GetBytes(signKey)),
                SecurityAlgorithms.HmacSha256)
        );
        return token;
    }

    private JwtSecurityToken? ValidateToken(string token)
    {
        var tokenHandler = new JwtSecurityTokenHandler();
        var signKey = _configuration.GetValue<string>("Jwt:Secret") ?? throw new ArgumentNullException();
        try
        {
            tokenHandler.ValidateToken(token, new TokenValidationParameters
            {
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(signKey)),
                ValidateIssuer = false,
                ValidateAudience = false
            }, out SecurityToken validatedToken);
            return (JwtSecurityToken)validatedToken;
        }
        catch (Exception)
        {
            return null;
        }
    }
}
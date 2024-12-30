using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using FluentResults;
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
            , audience: _configuration["Jwt:ValidAudience"], issuer: _configuration["Jwt:ValidIssuer"]
        );
        return token;
    }

    public Result<ClaimsPrincipal> ValidateToken(string token)
    {
        var tokenHandler = new JwtSecurityTokenHandler();
        var signKey = _configuration.GetValue<string>("Jwt:Secret") ?? throw new ArgumentNullException();
        try
        {
            var claimsPrincipal = tokenHandler.ValidateToken(token, new TokenValidationParameters
            {
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(signKey)),
                ValidateIssuer = false,
                ValidateAudience = false
            }, out _);
            return Result.Ok(claimsPrincipal);
        }
        catch (Exception e)
        {
            return Result.Fail(e.ToString());
        }
    }
}
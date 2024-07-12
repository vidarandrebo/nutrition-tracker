using System;
using System.IdentityModel.Tokens.Jwt;
using System.Threading.Tasks;
using FluentResults;
using NutritionTracker.Application.Identity;

namespace NutritionTracker.Application.Interfaces;

public interface IIdentityService
{

    public JwtSecurityToken CreateToken(Guid id, string email);
    public Task<Result<Guid>> RegisterUser(string email, string password);
    public Task<Result<ApplicationUserDto>> LoginUser(string email, string password);
}
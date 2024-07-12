using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Interfaces;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/[controller]")]
public class LoginController : ControllerBase
{
    private readonly IMediator _mediator;
    private readonly ILogger<LoginController> _logger;
    private readonly IIdentityService _identityService;

    public LoginController(IMediator mediator, ILogger<LoginController> logger, IIdentityService identityService)
    {
        _mediator = mediator;
        _logger = logger;
        _identityService = identityService;
    }
    // POST
    [HttpPost]
    public async Task<ActionResult<AccessTokenResponse>> Post(LoginRequest loginRequest)
    {
        var response = new AccessTokenResponse()
        {
            AccessToken = "",
            RefreshToken = "",
            ExpiresIn = 0,
        };
        var token = _identityService.CreateToken(Guid.NewGuid(), "test@mail.com");
        var tokenHandler = new JwtSecurityTokenHandler();
        var serializedToken = tokenHandler.WriteToken(token);
        Console.WriteLine(serializedToken);
        await Task.CompletedTask;
        return Created();
    }

}

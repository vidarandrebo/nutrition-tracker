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
[Route("api/auth/[controller]")]
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
        var response = await _identityService.LoginUser(loginRequest.Email, loginRequest.Password);
        if (response.IsFailed)
        {
            foreach(var err in response.Errors) {
                _logger.LogInformation(err.Message);
            }
            return BadRequest();
        }
        return Ok(response.Value);
    }

}

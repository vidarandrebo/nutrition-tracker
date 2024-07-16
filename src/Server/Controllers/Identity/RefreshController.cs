using System.Linq;
using System.Security.Claims;
using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Infrastructure.Identity;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/auth/[controller]")]
public class RefreshController : ControllerBase
{
    private readonly IMediator _mediator;
    private readonly ITokenHandler _tokenHandler;
    private readonly ILogger<LoginController> _logger;

    public RefreshController(IMediator mediator, ILogger<LoginController> logger, ITokenHandler tokenHandler)
    {
        _mediator = mediator;
        _logger = logger;
        _tokenHandler = tokenHandler;
    }

    // POST
    [HttpPost]
    public async Task<ActionResult<AccessTokenResponse>> Post(RefreshRequest refreshRequest)
    {
        _logger.LogInformation("refreshing token");

        var claimsResult = _tokenHandler.ValidateToken(refreshRequest.RefreshToken);

        if (claimsResult.IsFailed)
        {
            return Unauthorized();
        }

        var email = claimsResult.Value.Claims.FirstOrDefault(c => c.Type == ClaimTypes.Email);
        var id = claimsResult.Value.Claims.FirstOrDefault(c => c.Type == ClaimTypes.NameIdentifier);

        if (email is null || id is null)
        {
            return BadRequest();
        }

        var idGuid = await GuidHelper.GuidOrEmptyAsync(id.Value);

        var response = new AccessTokenResponse()
        {
            AccessToken = _tokenHandler.AccessToken(idGuid, email.Value),
            RefreshToken = _tokenHandler.RefreshToken(idGuid, email.Value),
            ExpiresIn = 60 * 60 * 24,
        };


        return Ok(response);
    }
}
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Threading.Tasks;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/[controller]")]
public class RefreshController : ControllerBase
{
    private readonly IMediator _mediator;

    private readonly ILogger<LoginController> _logger;

    public RefreshController(IMediator mediator, ILogger<LoginController> logger)
    {
        _mediator = mediator;
        _logger = logger;
    }
    // POST
    [HttpPost]
    public async Task<ActionResult<AccessTokenResponse>> Post(RefreshRequest refreshRequest)
    {
        _logger.LogInformation("refreshing token");
        var response = new AccessTokenResponse()
        {
            AccessToken = "",
            RefreshToken = "",
            ExpiresIn = 0,
        };
        await Task.CompletedTask;

        return Ok(response);
    }

}

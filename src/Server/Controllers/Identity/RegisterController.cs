using MediatR;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Threading.Tasks;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/auth/[controller]")]
public class RegisterController : ControllerBase
{
    private readonly IMediator _mediator;

    private readonly ILogger<LoginController> _logger;
    private readonly IIdentityService _identityService;

    public RegisterController(IMediator mediator, ILogger<LoginController> logger, IIdentityService identityService)
    {
        _mediator = mediator;
        _logger = logger;
        _identityService = identityService;
    }
    // POST
    [HttpPost]
    public async Task<ActionResult<HttpValidationProblemDetails>> Post(RegisterRequest registerRequest)
    {
        var result = await _identityService.RegisterUser(registerRequest.Email, registerRequest.Password);
        if (result.IsFailed) {
            foreach(var err in result.Errors) {
                _logger.LogInformation(err.Message);
            }
        }
        var response = new HttpValidationProblemDetails();
        await Task.CompletedTask;
        return response;
    }

}

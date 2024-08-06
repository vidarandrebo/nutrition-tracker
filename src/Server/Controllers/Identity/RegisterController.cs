using System.Threading;
using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Accounts;
using NutritionTracker.Application.Interfaces;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/auth/[controller]")]
public class RegisterController : ControllerBase
{
    private readonly IMediator _mediator;
    private readonly IConfiguration _cfg;
    private readonly ILogger<LoginController> _logger;
    private readonly IIdentityService _identityService;

    public RegisterController(IMediator mediator, ILogger<LoginController> logger, IIdentityService identityService, IConfiguration cfg)
    {
        _mediator = mediator;
        _logger = logger;
        _identityService = identityService;
        _cfg = cfg;
    }
    // POST
    [HttpPost]
    public async Task<ActionResult<HttpValidationProblemDetails>> Post(RegisterRequest registerRequest)
    {
<<<<<<< HEAD
        var result = await _identityService.RegisterUser(registerRequest.Email, registerRequest.Password);
        if (result.IsFailed) {
            foreach(var err in result.Errors) {
                _logger.LogInformation(err.Message);
            }
        }
=======
        var ctSource = new CancellationTokenSource(_cfg.GetValue<int>("CancellationToken:Delay"));
        var userId = await _identityService.RegisterUser(registerRequest.Email, registerRequest.Password);
        await _mediator.Send(
            new AddAccount.Request(userId.Value), ctSource.Token);
>>>>>>> 010a6ecc7c4e1aeeb31fbbffbec764f276364857
        var response = new HttpValidationProblemDetails();
        await Task.CompletedTask;
        return response;
    }

}

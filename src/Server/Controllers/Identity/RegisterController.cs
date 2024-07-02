using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Authentication.BearerToken;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity.Data;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace NutritionTracker.Server.Controllers.Identity;

[ApiController]
[Route("api/[controller]")]
public class RegisterController : ControllerBase
{
    private readonly IMediator _mediator;

    private readonly ILogger<LoginController> _logger;

    public RegisterController(IMediator mediator, ILogger<LoginController> logger)
    {
        _mediator = mediator;
        _logger = logger;
    }
    // POST
    [HttpPost]
    public async Task<ActionResult<HttpValidationProblemDetails>> Post(RegisterRequest registerRequest)
    {
        var response = new HttpValidationProblemDetails();
        await Task.CompletedTask;
        return response;
    } 
    
}
    
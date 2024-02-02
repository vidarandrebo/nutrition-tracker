using System.Threading.Tasks;
using MediatR;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace Server.Controllers;

[ApiController]
[Route("api/[controller]")]
public class FoodItemController : Controller
{
    private readonly IMediator _mediator;
    private readonly ILogger<FoodItemController> _logger;

    public FoodItemController(IMediator mediator, ILogger<FoodItemController> logger)
    {
        _mediator = mediator;
        _logger = logger;
    }

    [HttpGet]
    public async Task<ActionResult> Get()
    {
        var getUserIdResult = HttpContext.GetUserName();
        if (getUserIdResult.IsSuccess)
        {
            _logger.LogInformation(getUserIdResult.Value.ToString());
        }
        else
        {
            _logger.LogError(getUserIdResult.Errors.ToString());
        }
        return Ok();
    }
}
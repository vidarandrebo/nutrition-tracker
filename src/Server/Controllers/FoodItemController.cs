using MediatR;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.FoodItems;
using NutritionTracker.Domain.FoodItems.Contracts;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;

namespace NutritionTracker.Server.Controllers;

[Authorize]
[ApiController]
[Route("api/[controller]")]
public class FoodItemController : ControllerBase
{
    private readonly IMediator _mediator;
    private readonly ILogger<FoodItemController> _logger;

    public FoodItemController(IMediator mediator, ILogger<FoodItemController> logger)
    {
        _mediator = mediator;
        _logger = logger;
    }

    [HttpGet]
    public async Task<ActionResult<FoodItemResponse[]>> Get()
    {
        var ctSrc = new CancellationTokenSource(2000);
        var getFoodItemRequest = new GetFoodItems.Request();
        var getFoodItemResult = await _mediator.Send(getFoodItemRequest, ctSrc.Token);
        if (getFoodItemResult.IsFailed)
        {
            return BadRequest(getFoodItemResult.Errors);
        }
        var response = FoodItemResponse.FromDtos(getFoodItemResult.Value);

        return Ok(getFoodItemResult.Value);
    }


    [HttpPost]
    public async Task<ActionResult<FoodItemResponse>> PostAsync(PostFoodItemRequest form)
    {
        var getUserIdResult = HttpContext.GetUserId();
        if (getUserIdResult.IsFailed)
        {
            // Should not happen as user is authorized.
            return BadRequest();
        }

        var createFoodItemRequest = new AddFoodItem.Request(form, getUserIdResult.Value);
        using var ctSrc = new CancellationTokenSource(2000);
        var createFoodItemResult = await _mediator.Send(createFoodItemRequest, ctSrc.Token);
        if (createFoodItemResult.IsFailed)
        {
            return BadRequest(createFoodItemResult.Errors);
        }
        var response = FoodItemResponse.FromDto(createFoodItemResult.Value);

        return Created(nameof(Get), response);
    }
}
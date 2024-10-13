using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Diary.Entities;

namespace NutritionTracker.Web.Pages.Diary;

public partial class Index
{
    [Inject] private ILogger<Index> _logger { get; init; }
    [Inject] private IApplicationDbContext _db { get; init; }
    public Day Day { get; set; }
    private DateOnly _selectedDate = DateOnly.FromDateTime(DateTime.Now);

    public void Hello()
    {
        _logger.LogInformation("hello there");
    }

    public DateOnly SelectedDate
    {
        get { return _selectedDate; }
        set
        {
            _selectedDate = value;
            _logger.LogInformation("Set date to {date}", _selectedDate);
        }
    }

    protected override async Task OnInitializedAsync()
    {
        var day = await _db.Days.FirstOrDefaultAsync();
        if (day is not null)
        {
            Day = day;
        }
        else
        {
            Day = new Day();
        }

        await base.OnInitializedAsync();
    }
}
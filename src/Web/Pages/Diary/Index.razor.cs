using System;
using System.Linq;
using System.Linq.Expressions;
using System.Runtime.InteropServices.JavaScript;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Forms;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using NutritionTracker.Application.Interfaces;
using NutritionTracker.Domain.Diary.Entities;

namespace NutritionTracker.Web.Pages.Diary;

[Authorize]
public partial class Index
{
    [Inject] private ILogger<Index> Logger { get; init; }
    [Inject] private IApplicationDbContext Db { get; init; }
    public Day Day { get; set; }

    public async Task AddMeal()
    {
        Logger.LogInformation("hello there");
        var ctSrc = new CancellationTokenSource(5000);
        var meal = new Meal();
        Day.AddMeal(meal);
        await Db.SaveChangesAsync(ctSrc.Token);
        ;
    }

    public async Task DecrementDay()
    {
        var date = SelectedDate.AddDays(-1);
        await SetDate(date);
    }
    public async Task IncrementDay()
    {
        var date = SelectedDate.AddDays(1);
        await SetDate(date);
    }
    public async Task SetDate(DateOnly date)
    {
        Logger.LogError("hello from value changed");
        SelectedDate = date;
        Logger.LogInformation("Set date to {date}", SelectedDate);
        await GetDay();
    }

    public DateOnly SelectedDate { get; set; } = DateOnly.FromDateTime(DateTime.Now);

    public async Task GetDay()
    {
        var ctSrc = new CancellationTokenSource(5000);
        var day = await Db.Days
            .Where(d => d.Date == SelectedDate)
            .Include(d => d.Meals)
            .FirstOrDefaultAsync(ctSrc.Token);
        if (day is not null)
        {
            Day = day;
        }
        else
        {
            Day = new Day();
            Day.Date = SelectedDate;
            Db.Days.Add(Day);
            await Db.SaveChangesAsync(ctSrc.Token);
        }
    }

    protected override async Task OnInitializedAsync()
    {
        await GetDay();
        await base.OnInitializedAsync();
    }
}
using NutritionTracker.Infrastructure;
using Microsoft.EntityFrameworkCore;
using Microsoft.Data.Sqlite;

namespace NutritionTracker.Application.Tests;

public static class DatabaseHelper
{
    public static ApplicationDbContext NewContext()
    {
        var conn = new SqliteConnection("Filename=:memory:");
        conn.Open();

        var contextOptions = new DbContextOptionsBuilder<ApplicationDbContext>()
            .UseSqlite(conn, options => { options.MigrationsAssembly("NutritionTracker.Migrations.Sqlite"); })
            .Options;

        var ctx = new ApplicationDbContext(contextOptions);
        ctx.Database.Migrate();

        return ctx;
    }
}
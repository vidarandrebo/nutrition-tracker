using Infrastructure;
using Microsoft.EntityFrameworkCore;
using Microsoft.Data.Sqlite;

namespace Application.Tests;

public static class DatabaseHelper
{
    public static ApplicationDbContext NewContext()
    {
        var conn = new SqliteConnection("Filename=:memory:");
        conn.Open();

        var contextOptions = new DbContextOptionsBuilder<ApplicationDbContext>()
            .UseSqlite(conn, options => {
                options.MigrationsAssembly(Migrations.Sqlite.Provider.Assembly);
            })
            .Options;

        var ctx = new ApplicationDbContext(contextOptions);
        ctx.Database.Migrate();

        return ctx;
    }
}
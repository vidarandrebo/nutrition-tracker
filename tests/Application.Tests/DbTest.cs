using Application.Interfaces;
using Infrastructure;

namespace Application.Tests;

public class DbTest
{
    private readonly IApplicationDbContext _db;

    public DbTest()
    {
        _db = DatabaseHelper.NewContext();
    }

    [Fact]
    public void NotNullTest()
    {
        Assert.NotNull(_db);
    }
}
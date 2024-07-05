using System.Data;
using System.Globalization;

namespace NutritionTracker.DataImporter;

public static class Program
{
    public static void Main(string[] args)
    {
        var date = DateOnly.Parse("2022-03-10");
        Console.WriteLine(date.ToString("O", CultureInfo.InvariantCulture));
    }
}
using System.Data;
using System.Globalization;
using Newtonsoft.Json;
using NutritionTracker.DataImporter.Models.Matvareportalen;
using NutritionTracker.Domain.FoodItems.Entities;

namespace NutritionTracker.DataImporter;

public static class Program
{
    public static void Main(string[] args)
    {
        var date = DateOnly.Parse("2022-03-10");
        Console.WriteLine(date.ToString("O", CultureInfo.InvariantCulture));

        var json = File.ReadAllText(Path.Join("Data", "Matvareportalen", "foods.json"));
        var foodItems = JsonConvert.DeserializeObject<Root>(json);

        Console.WriteLine($"Found {foodItems.Foods.Count} items");
        foreach(var item in foodItems.Foods) {
            Console.WriteLine(item.FoodName);
            foreach(var constituent in item.Constituents) {
                Console.WriteLine($"{constituent.NutrientId}: {constituent.Quantity}{constituent.Unit}");
            }
            Console.Write("Add to storage Y/n");
            var input = Console.ReadLine();
            switch (input) {
                case "n":
                case "N":
                    Console.WriteLine("discard");
                    break;
                case "":
                case "Y":
                case "y":
                    Console.WriteLine("adding");
                    break;
                default:
                    continue;
            }
        }
    }
}
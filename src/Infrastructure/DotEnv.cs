using System;
using System.Globalization;
using System.IO;
using Microsoft.Extensions.Configuration;

namespace NutritionTracker.Infrastructure;

public static class DotEnv
{
    public static void LoadEnvToConfiguration(this ConfigurationManager cfg, string filePath)
    {
        if (!File.Exists(filePath))
        {
            Console.WriteLine("file does not exist");
            return;
        }

        foreach (var line in File.ReadAllLines(filePath))
        {
            var parts = line.Split(
                '=',
                StringSplitOptions.RemoveEmptyEntries);

            if (parts.Length != 2)
                continue;

            var key = parts[0];
            var value = parts[1];

            key = key.ToLower().Replace("_", " ");
            var info = CultureInfo.CurrentCulture.TextInfo;
            key = info.ToTitleCase(key).Replace(" ", ":");

            cfg[key] = value;
        }
    }
}
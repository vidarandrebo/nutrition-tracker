using Microsoft.EntityFrameworkCore;

namespace Domain;

[Owned]
public class NutritionalContent
{
    public double Protein { get; set; }
    public double Carbohydrate { get; set; }
    public double Fat { get; set; }
    public double Kcal { get; set; }


    public NutritionalContent(double protein, double carbohydrate, double fat, double kcal)
    {
        Protein = protein;
        Carbohydrate = carbohydrate;
        Fat = fat;
        Kcal = kcal;
    }
}
using Microsoft.EntityFrameworkCore;

namespace Domain;

[Owned]
public class NutritionalContent
{
    public double Protein;
    public double Carbohydrate;
    public double Fat;
    public double Kcal;

    public NutritionalContent()
    {
    }

    public NutritionalContent(double protein, double carbohydrate, double fat, double kcal)
    {
        Protein = protein;
        Carbohydrate = carbohydrate;
        Fat = fat;
        Kcal = kcal;
    }
}
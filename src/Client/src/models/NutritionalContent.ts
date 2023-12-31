/**
 * Represents the nutritional content in 100 grams of the food item.
 */
export class NutritionalContent {
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    unit: "grams" | "ml";

    constructor(protein: number, carbohydrate: number, fat: number, kCal: number, unit: "grams"|"ml") {
        this.protein = protein;
        this.carbohydrate = carbohydrate;
        this.fat = fat;
        this.kCal = kCal;
        this.unit = unit;
    }
}
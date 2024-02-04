import {IObjectAssignable} from "./ObjectAssignable.ts";

/**
 * Represents the nutritional content in 100 grams of the food item.
 */
export class NutritionalContent implements IObjectAssignable {
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    unit: "grams" | "ml";

    constructor(protein: number, carbohydrate: number, fat: number, kCal: number, unit: "grams" | "ml") {
        this.protein = protein;
        this.carbohydrate = carbohydrate;
        this.fat = fat;
        this.kCal = kCal;
        this.unit = unit;
    }

    assignFromObject(src: Record<string, never>): void {
        if (Object.prototype.hasOwnProperty.call(src, "protein")) {
            this.protein = src["protein"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "carbohydrate")) {
            this.carbohydrate = src["carbohydrate"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "fat")) {
            this.fat = src["fat"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "kCal")) {
            this.kCal = src["kCal"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "unit")) {
            this.unit = src["unit"]
        }
    }
}
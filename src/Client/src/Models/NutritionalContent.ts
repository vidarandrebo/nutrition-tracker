import { ObjectAssignable } from "./ObjectAssignable.ts";

/**
 * Represents the nutritional content in 100 grams of the food item.
 */
export class NutritionalContent extends ObjectAssignable {
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    unit: "grams" | "ml";

    constructor() {
        super();
        this.protein = 0.0;
        this.carbohydrate = 0.0;
        this.fat = 0.0;
        this.kCal = 0.0;
        this.unit = "grams";
    }
}

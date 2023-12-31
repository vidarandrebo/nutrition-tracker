import {NutritionalContent} from "./NutritionalContent.ts";
import {FoodItemForm} from "./FoodItemForm.ts";
import {v4 as UuidV4} from "uuid";

export class FoodItem {
    id: string;
    brand: string;
    productName: string;
    nutritionalContent: NutritionalContent;
    ownerId: string;

    constructor(id: string, brand: string, productName: string, nutritionalContent: NutritionalContent, ownerId: string) {
        this.id = id;
        this.brand = brand;
        this.productName = productName;
        this.nutritionalContent = nutritionalContent;
        this.ownerId = ownerId;
    }
}

export function postFoodItem(foodForm: FoodItemForm): FoodItem {
    // post stuff goes here
    const fid = UuidV4();
    const uid = UuidV4();
    return new FoodItem(
        fid,
        foodForm.brand,
        foodForm.productName,
        new NutritionalContent(
            foodForm.protein,
            foodForm.carbohydrate,
            foodForm.fat,
            foodForm.kCal,
            foodForm.unit),
        uid);
}
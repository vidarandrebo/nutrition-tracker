import {NutritionalContent} from "./NutritionalContent.ts";

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
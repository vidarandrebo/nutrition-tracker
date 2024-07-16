import { Macronutrients } from "./Macronutrients.ts";
import { FoodItemForm } from "./FoodItemForm.ts";
import { HttpRequest } from "http-methods-ts";
import { loadUser } from "./User.ts";
import { IObjectAssignable } from "./ObjectAssignable.ts";

export class FoodItem implements IObjectAssignable {
    id: string;
    brand: string;
    productName: string;
    macronutrients: Macronutrients;
    ownerId: string;

    constructor() {
        this.id = "";
        this.brand = "";
        this.productName = "";
        this.ownerId = "";
        this.macronutrients = new Macronutrients();
    }

    assignFromObject(src: Record<string, never>): void {
        if (Object.prototype.hasOwnProperty.call(src, "id")) {
            this.id = src["id"];
        }
        if (Object.prototype.hasOwnProperty.call(src, "brand")) {
            this.brand = src["brand"];
        }
        if (Object.prototype.hasOwnProperty.call(src, "productName")) {
            this.productName = src["productName"];
        }
        if (Object.prototype.hasOwnProperty.call(src, "ownerId")) {
            this.ownerId = src["ownerId"];
        }
        if (Object.prototype.hasOwnProperty.call(src, "macronutrients")) {
            this.macronutrients.assignFromObject(src["macronutrients"]);
        }
    }
}

export async function postFoodItem(foodForm: FoodItemForm): Promise<FoodItem | null> {
    // send post request to server
    const user = loadUser();
    if (user == null) {
        return null;
    }
    const httpRequest = new HttpRequest()
        .setRoute("api/fooditem")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(foodForm)
        .setBearerToken(user.accessToken);
    await httpRequest.send();
    const response = httpRequest.getResponseData();
    if (response != null && response.body != null) {
        return fromObject(response.body as Record<string, never>);
    }
    return null;
}

export async function getFoodItems(): Promise<FoodItem[]> {
    const user = loadUser();
    if (user == null) {
        return [];
    }
    const httpRequest = new HttpRequest()
        .setRoute("api/fooditem")
        .setMethod("GET")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken);
    await httpRequest.send();
    const response = httpRequest.getResponseData();
    if (response == null || response.body == null) {
        return [];
    }
    const foodItems = response.body as Record<string, never>[];
    return foodItems.map((x) => fromObject(x));
}

function fromObject(src: Record<string, never>): FoodItem {
    const item = new FoodItem();
    item.assignFromObject(src);
    return item;
}

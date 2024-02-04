import {NutritionalContent} from "./NutritionalContent.ts";
import {FoodItemForm} from "./FoodItemForm.ts";
import {HttpRequest} from "./Http.ts";
import {loadUser} from "./User.ts";
import {IObjectAssignable} from "./ObjectAssignable.ts";

export class FoodItem implements IObjectAssignable {
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

    assignFromObject(src: Record<string, never>): void {
        if (Object.prototype.hasOwnProperty.call(src, "id")) {
            this.id = src["id"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "brand")) {
            this.brand = src["brand"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "productName")) {
            this.productName = src["productName"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "ownerId")) {
            this.id = src["ownerId"]
        }
        if (Object.prototype.hasOwnProperty.call(src, "nutritionalContent")) {
            this.nutritionalContent.assignFromObject(src["nutritionalContent"])
        }
    }
}


export async function postFoodItem(foodForm: FoodItemForm): Promise<FoodItem | null> {
    // send post request to server
    const user = loadUser();
    if (user == null) {
        return null
    }
    const httpRequest = new HttpRequest()
        .setRoute("api/fooditem")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(foodForm)
        .setBearerToken(user.accessToken);
    await httpRequest.send();
    const response = httpRequest.getResponseData();
    if (response != null)
        console.log(response)
    return null
}

export async function getFoodItems(): Promise<FoodItem[]> {
    const user = loadUser();
    if (user == null) {
        return []
    }
    const httpRequest = new HttpRequest()
        .setRoute("api/fooditem")
        .setMethod("GET")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken);
    await httpRequest.send();
    const response = httpRequest.getResponseData();
    if (response == null) {
        return [];
    }
    return [];
}
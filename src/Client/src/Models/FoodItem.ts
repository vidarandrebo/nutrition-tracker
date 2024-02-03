import {NutritionalContent} from "./NutritionalContent.ts";
import {FoodItemForm} from "./FoodItemForm.ts";
import {v4 as UuidV4} from "uuid";
import {HttpRequest} from "./Http.ts";
import {loadUser} from "./User.ts";

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

export async function postFoodItem(foodForm: FoodItemForm): Promise<FoodItem> {
    // TODO
    // send post request to server
    const fid = UuidV4(); //TMP
    const uid = UuidV4(); //TMP
    const user = loadUser();
    if (user == null) {
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
    const httpRequest = new HttpRequest()
        .setRoute("api/fooditem")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(foodForm)
        .setBearerToken(user.accessToken);
    await httpRequest.send();
    const response = httpRequest.getResponseData();
    console.log(response)
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
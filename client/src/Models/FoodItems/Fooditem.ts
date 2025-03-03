import { HttpRequest } from "http-methods-ts";
import { readFromLocalStorage } from "../User.ts";

export class FoodItem {
    id: string;
    manufacturer: string;
    product: string;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number
    source: string;

    constructor() {
        this.id = "";
        this.manufacturer = "";
        this.product = "";
        this.protein = 0.0;
        this.carbohydrate = 0.0;
        this.fat = 0.0;
        this.kCal = 0.0;
        this.source = "";
    }

    static assignFromObject(obj: Record<string, never>): FoodItem {
        const foodItem = new FoodItem();
        foodItem.id = obj["id"];
        foodItem.manufacturer = obj["manufacturer"];
        foodItem.product = obj["product"];
        foodItem.protein = obj["protein"];
        foodItem.carbohydrate = obj["carbohydrate"];
        foodItem.fat = obj["fat"];
        foodItem.kCal = obj["kCal"]
        foodItem.source = obj["source"];
        return foodItem;
    }
}

export async function getFoodItems(): Promise<FoodItem[] | null> {
    const user = readFromLocalStorage();
    if (user === null) {
        return null;
    }
    const request = new HttpRequest()
        .setRoute("/api/food-items")
        .setMethod("GET")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken);
    await request.send();
    const response = request.getResponseData();
    if (response === null) {
        return null;
    }
    if (response.status === 200) {
        const payload = response.body as [];
        const foodItems = payload.map((item) => FoodItem.assignFromObject(item));
        return foodItems;
    }
    return null;
}

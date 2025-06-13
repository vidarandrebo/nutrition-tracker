import { HttpRequest } from "http-methods-ts";
import type { FoodItemResponse } from "./Responses.ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import type { Energy } from "../Common/Energy.ts";

export class FoodItem {
    id: number;
    manufacturer: string;
    product: string;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    source: string;
    get name(): string {
        let out = this.manufacturer;
        if (this.product.length > 0 && out.length > 0) {
            out += " ";
        }
        return out + this.product;
    }
    get energy(): Energy {
        return {
            protein: this.protein,
            carbohydrate: this.carbohydrate,
            fat: this.fat,
            kCal: this.kCal,
        };
    }

    constructor() {
        this.id = 0;
        this.manufacturer = "";
        this.product = "";
        this.protein = 0.0;
        this.carbohydrate = 0.0;
        this.fat = 0.0;
        this.kCal = 0.0;
        this.source = "";
    }

    static fromResponse(res: FoodItemResponse): FoodItem {
        const foodItem = new FoodItem();
        foodItem.id = res.id;
        foodItem.manufacturer = res.manufacturer;
        foodItem.product = res.product;
        foodItem.protein = res.protein;
        foodItem.carbohydrate = res.carbohydrate;
        foodItem.fat = res.fat;
        foodItem.kCal = res.kCal;
        foodItem.source = res.source;
        return foodItem;
    }

    static async get(): Promise<FoodItem[] | null> {
        const userStore = useUserStore();
        const user = userStore.user;
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
            const payload = response.body as FoodItemResponse[];
            const foodItems = payload.map((item) => FoodItem.fromResponse(item));
            return foodItems;
        }
        return null;
    }
    static async getById(id: number): Promise<FoodItem | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return null;
        }
        const request = new HttpRequest()
            .setRoute(`/api/food-items/${id}`)
            .setMethod("GET")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken);
        await request.send();
        const response = request.getResponseData();
        if (response === null) {
            return null;
        }
        if (response.status === 200) {
            const payload = response.body as FoodItemResponse;
            return FoodItem.fromResponse(payload);
        }
        return null;
    }
}

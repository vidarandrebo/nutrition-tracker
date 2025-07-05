import { HttpRequest } from "http-methods-ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import type { Energy } from "../Common/Energy.ts";
import { type FoodItemResponse } from "../../Gen";
import { getFoodItemsClient } from "../Api.ts";

export class FoodItem {
    id: number;
    manufacturer: string;
    product: string;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    public: boolean;
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
        this.public = false;
        this.source = "";
    }

    static fromResponse(res: FoodItemResponse): FoodItem {
        const foodItem = new FoodItem();
        foodItem.id = res.id ?? 0;
        foodItem.manufacturer = res.manufacturer ?? "";
        foodItem.product = res.product ?? "";
        foodItem.protein = res.protein ?? 0.0;
        foodItem.carbohydrate = res.carbohydrate ?? 0.0;
        foodItem.fat = res.fat ?? 0.0;
        foodItem.kCal = res.kCal ?? 0.0;
        foodItem.public = res.isPublic ?? false;
        foodItem.source = res.source ?? "";
        return foodItem;
    }

    static async get(): Promise<FoodItem[] | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return null;
        }
        const api = getFoodItemsClient();
        const response = await api.apiFoodItemsGet();

        const foodItems = response.map((item) => FoodItem.fromResponse(item));
        return foodItems;
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

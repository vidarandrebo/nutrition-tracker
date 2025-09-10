import type { Energy } from "../Common/Energy.ts";
import { type FoodItemResponse } from "../../Gen";
import { getFoodItemsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import { PortionSize } from "./PortionSize.ts";

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
    ownerId: number;
    portionSizes: PortionSize[];

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
        this.ownerId = 0;
        this.portionSizes = [];
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
        foodItem.public = res.isPublic;
        foodItem.source = res.source;
        foodItem.ownerId = res.ownerId;
        foodItem.portionSizes = PortionSize.fromResponses(res.portionSizes ?? []);
        return foodItem;
    }

    static async get(): Promise<FoodItem[] | null> {
        const client = getFoodItemsClient();
        try {
            const response = await client.apiFoodItemsGet();
            const foodItems = response.map((item) => FoodItem.fromResponse(item));
            return foodItems;
        } catch {
            console.log("failed to fetch food items");
        }
        return null;
    }

    static async getById(id: number): Promise<FoodItem | null> {
        const client = getFoodItemsClient();
        try {
            const response = await client.apiFoodItemsIdGet({ id: id });
            return FoodItem.fromResponse(response);
        } catch {
            console.log("failed to fetch food items");
        }
        return null;
    }
    static async delete(id: number): Promise<Result<void>> {
        const client = getFoodItemsClient();

        return await tryCatch(
            client.apiFoodItemsIdDelete({
                id: id,
            }),
        );
    }
}

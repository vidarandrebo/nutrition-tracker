import type { PostMealEntryRequest } from "./Requests.ts";
import { getMealsClient } from "../Api.ts";
import type { MealEntryResponse } from "../../Gen";

export class MealEntry {
    id: number;
    foodItemId: number | null;
    recipeId: number | null;
    amount: number;

    constructor() {
        this.id = 0;
        this.foodItemId = null;
        this.recipeId = null;
        this.amount = 0;
    }

    static async add(entry: PostMealEntryRequest, mealId: number): Promise<MealEntry | null> {
        const client = getMealsClient();

        try {
            const response = await client.apiMealsIdEntriesPost({ id: mealId, postMealEntryRequest: entry });
            return MealEntry.fromResponse(response);
        } catch {
            console.log("failed to add entry to meal");
        }
        return null;
    }

    static fromResponses(res: MealEntryResponse[]): MealEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealEntryResponse): MealEntry {
        const me = new MealEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.foodItemId = res.foodItemId ?? null;
        me.recipeId = res.recipeId ?? null;
        return me;
    }
}

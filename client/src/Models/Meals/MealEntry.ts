import type { PostMealEntryRequest } from "./Requests.ts";
import { getMealsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import type { MealFoodItemEntryResponse } from "../../Gen";

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
            const response = await client.apiMealsIdFoodItemEntriesPost({
                id: mealId,
                mealFoodItemEntryPostRequest: entry,
            });
            return MealEntry.fromResponse(response);
        } catch {
            console.log("failed to add entry to meal");
        }
        return null;
    }
    static async delete(id: number, mealId: number): Promise<Result<void>> {
        const client = getMealsClient();

        return await tryCatch(
            client.apiMealsMealIdEntriesEntryIdDelete({
                entryId: id,
                mealId: mealId,
            }),
        );
    }

    static fromResponses(res: MealFoodItemEntryResponse[]): MealEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealFoodItemEntryResponse): MealEntry {
        const me = new MealEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.foodItemId = res.foodItemId;
        return me;
    }
}

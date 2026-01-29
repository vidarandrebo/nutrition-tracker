import { getMealsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import type { MealFoodItemEntryPostRequest, MealFoodItemEntryResponse } from "../../Gen";

export class MealFoodItemEntry {
    id: number;
    foodItemId: number;
    amount: number;

    constructor() {
        this.id = 0;
        this.foodItemId = 0;
        this.amount = 0;
    }

    static async add(entry: MealFoodItemEntryPostRequest, mealId: number): Promise<MealFoodItemEntry | null> {
        const client = getMealsClient();

        try {
            const response = await client.apiMealsMealIdFoodItemEntriesPost({
                mealId: mealId,
                mealFoodItemEntryPostRequest: entry,
            });
            return MealFoodItemEntry.fromResponse(response);
        } catch {
            console.log("failed to add entry to meal");
        }
        return null;
    }

    static async delete(id: number, mealId: number): Promise<Result<void>> {
        const client = getMealsClient();

        return await tryCatch(
            client.apiMealsMealIdFoodItemEntriesFoodItemEntryIdDelete({
                foodItemEntryId: id,
                mealId: mealId,
            }),
        );
    }

    static fromResponses(res: MealFoodItemEntryResponse[]): MealFoodItemEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealFoodItemEntryResponse): MealFoodItemEntry {
        const me = new MealFoodItemEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.foodItemId = res.foodItemId;
        return me;
    }
}

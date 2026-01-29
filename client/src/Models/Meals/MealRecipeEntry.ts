import { getMealsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import type { MealRecipeEntryPostRequest, MealRecipeEntryResponse } from "../../Gen";

export class MealRecipeEntry {
    id: number;
    recipeId: number;
    amount: number;

    constructor() {
        this.id = 0;
        this.recipeId = 0;
        this.amount = 0;
    }

    static async add(entry: MealRecipeEntryPostRequest, mealId: number): Promise<MealRecipeEntry | null> {
        const client = getMealsClient();

        try {
            const response = await client.apiMealsMealIdRecipeEntriesPost({
                mealId: mealId,
                mealRecipeEntryPostRequest: entry,
            });
            return MealRecipeEntry.fromResponse(response);
        } catch {
            console.log("failed to add entry to meal");
        }
        return null;
    }
    static async delete(id: number, mealId: number): Promise<Result<void>> {
        const client = getMealsClient();

        return await tryCatch(
            client.apiMealsMealIdRecipeEntriesRecipeEntryIdDelete({
                recipeEntryId: id,
                mealId: mealId,
            }),
        );
    }

    static fromResponses(res: MealRecipeEntryResponse[]): MealRecipeEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealRecipeEntryResponse): MealRecipeEntry {
        const me = new MealRecipeEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.recipeId = res.recipeId;
        return me;
    }
}

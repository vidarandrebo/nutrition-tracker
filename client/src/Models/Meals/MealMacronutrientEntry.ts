import { getMealsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import type { MealMacronutrientEntryPostRequest, MealMacronutrientEntryResponse } from "../../Gen";

export class MealMacronutrientEntry {
    id: number;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;

    constructor() {
        this.id = 0;
        this.protein = 0;
        this.carbohydrate = 0;
        this.fat = 0;
        this.kCal = 0;
    }

    static async add(entry: MealMacronutrientEntryPostRequest, mealId: number): Promise<MealMacronutrientEntry | null> {
        const client = getMealsClient();

        try {
            const response = await client.apiMealsMealIdMacronutrientEntriesPost({
                mealId: mealId,
                mealMacronutrientEntryPostRequest: entry,
            });
            return MealMacronutrientEntry.fromResponse(response);
        } catch {
            console.log("failed to add entry to meal");
        }
        return null;
    }

    static async delete(id: number, mealId: number): Promise<Result<void>> {
        const client = getMealsClient();

        return await tryCatch(
            client.apiMealsMealIdMacronutrientEntriesMacronutrientEntryIdDelete({
                macronutrientEntryId: id,
                mealId: mealId,
            }),
        );
    }

    static fromResponses(res: MealMacronutrientEntryResponse[]): MealMacronutrientEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealMacronutrientEntryResponse): MealMacronutrientEntry {
        const me = new MealMacronutrientEntry();
        me.id = res.id;
        me.protein = res.protein;
        me.carbohydrate = res.carbohydrate;
        me.fat = res.fat;
        me.kCal = res.kCal;
        return me;
    }
}

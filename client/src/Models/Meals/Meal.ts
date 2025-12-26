import { MealRecipeEntry } from "./MealRecipeEntry.ts";
import { addDays, isToday, startOfDay } from "../../Utilities/Date.ts";
import { getMealsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import { MealFoodItemEntry } from "./MealFoodItemEntry.ts";
import type { MealPostRequest, MealResponse } from "../../Gen";
import type { MealMacronutrientEntry } from "./MealMacronutrientEntry.ts";

export class Meal {
    id: number;
    timestamp: Date;
    sequenceNumber: number;
    foodItemEntries: MealFoodItemEntry[];
    recipeEntries: MealRecipeEntry[];
    macronutrientEntries: MealMacronutrientEntry[];

    constructor() {
        this.id = 0;
        this.timestamp = new Date();
        this.sequenceNumber = 0;
        this.foodItemEntries = [];
        this.recipeEntries = [];
        this.macronutrientEntries = [];
    }

    static mealTimeStamp(day: Date): Date {
        const ts = isToday(day) ? new Date() : startOfDay(day);
        if (isNaN(ts.getUTCSeconds())) {
            return new Date();
        }
        return ts;
    }

    static async add(day: Date): Promise<Meal | null> {
        const request: MealPostRequest = {
            timestamp: this.mealTimeStamp(day),
        };

        const client = getMealsClient();
        try {
            const response = await client.apiMealsPost({ mealPostRequest: request });
            return Meal.fromResponse(response);
        } catch {
            console.log("oi, ya goofed up");
        }
        return null;
    }
    static async delete(id: number): Promise<Result<void>> {
        const client = getMealsClient();

        return await tryCatch(
            client.apiMealsIdDelete({
                id: id,
            }),
        );
    }

    static async getByDay(day: Date): Promise<Meal[] | null> {
        const dateFrom = startOfDay(day);
        const dateTo = addDays(dateFrom, 1);
        const client = getMealsClient();
        try {
            const response = await client.apiMealsGet({ dateFrom: dateFrom, dateTo: dateTo });
            return Meal.fromResponses(response);
        } catch {
            console.warn("failed to fetch meals from server");
        }
        return null;
    }

    static async getById(id: number): Promise<Meal | null> {
        const client = getMealsClient();
        try {
            const response = await client.apiMealsIdGet({ id: id });
            return Meal.fromResponse(response);
        } catch {
            console.warn("failed to fetch meals from server");
        }
        return null;
    }

    static fromResponse(res: MealResponse): Meal {
        const m = new Meal();
        m.id = res.id;
        m.timestamp = new Date(res.timestamp);
        m.sequenceNumber = res.sequenceNumber;
        m.recipeEntries = MealRecipeEntry.fromResponses(res.recipeEntries ?? []);
        m.foodItemEntries = MealFoodItemEntry.fromResponses(res.foodItemEntries ?? []);
        return m;
    }

    static fromResponses(res: MealResponse[]): Meal[] {
        return res.map((r) => this.fromResponse(r));
    }
}

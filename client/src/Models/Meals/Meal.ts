import { MealEntry } from "./MealEntry.ts";
import { addDays, isToday, startOfDay } from "../../Utilities/Date.ts";
import type { PostMealRequest } from "./Requests.ts";
import type { MealResponse } from "../../Gen";
import { getMealsClient } from "../Api.ts";

export class Meal {
    id: number;
    timestamp: Date;
    sequenceNumber: number;
    entries: MealEntry[];

    constructor() {
        this.id = 0;
        this.timestamp = new Date();
        this.sequenceNumber = 0;
        this.entries = [];
    }

    static mealTimeStamp(day: Date): Date {
        const ts = isToday(day) ? new Date() : startOfDay(day);
        if (isNaN(ts.getUTCSeconds())) {
            return new Date();
        }
        return ts;
    }
    static async add(day: Date): Promise<Meal | null> {
        const request: PostMealRequest = {
            timestamp: this.mealTimeStamp(day),
        };

        const client = getMealsClient();
        try {
            const response = await client.apiMealsPost({ postMealRequest: request });
            return Meal.fromResponse(response);
        } catch {
            console.log("oi, ya goofed up");
        }
        return null;
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
        m.id = res.id ?? 0;
        m.timestamp = new Date(res.timestamp ?? "");
        m.sequenceNumber = res.sequenceNumber ?? 0;
        m.entries = MealEntry.fromResponses(res.entries ?? []);
        return m;
    }

    static fromResponses(res: MealResponse[]): Meal[] {
        return res.map((r) => this.fromResponse(r));
    }
}

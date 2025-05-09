import type { MealEntryResponse } from "./Responses.ts";

export class MealEntry {
    id: number;
    foodItemId: number;
    amount: number;

    constructor() {
        this.id = 0;
        this.foodItemId = 0;
        this.amount = 0;
    }

    static fromResponses(res: MealEntryResponse[]): MealEntry[] {
        return res.map(v => this.fromResponse(v));
    }

    static fromResponse(res: MealEntryResponse): MealEntry {
        const me = new MealEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.foodItemId = res.foodItemId;
        return me;
    }
}
import type { RecipeFoodItemEntryEntryResponse } from "../../Gen";

export class RecipeFoodItemEntry {
    id: number;
    amount: number;
    foodItemId: number;

    constructor() {
        this.id = 0;
        this.amount = 0;
        this.foodItemId = 0;
    }

    static fromResponse(request: RecipeFoodItemEntryEntryResponse): RecipeFoodItemEntry {
        const r = new RecipeFoodItemEntry();
        r.id = request.id;
        r.amount = request.amount;
        r.foodItemId = request.foodItemId;
        return r;
    }
}

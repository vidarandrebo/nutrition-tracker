import type { RecipeEntryResponse } from "../../Gen";

export class RecipeEntry {
    id: number;
    amount: number;
    foodItemId: number;

    constructor() {
        this.id = 0;
        this.amount = 0;
        this.foodItemId = 0;
    }

    static fromResponse(request: RecipeEntryResponse): RecipeEntry {
        const r = new RecipeEntry();
        r.id = request.id;
        r.amount = request.amount;
        r.foodItemId = request.foodItemId;
        return r;
    }
}

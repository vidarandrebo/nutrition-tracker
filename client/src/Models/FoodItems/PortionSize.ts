import { getFoodItemsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";
import type { FoodItemPortionSizeResponse } from "../../Gen";

export class PortionSize {
    id: number;
    amount: number;
    name: string;
    constructor() {
        this.id = 0;
        this.amount = 0;
        this.name = "";
    }
    static fromResponse(res: FoodItemPortionSizeResponse): PortionSize {
        const p = new PortionSize();
        p.id = res.id;
        p.amount = res.amount;
        p.name = res.name;
        return p;
    }
    static fromResponses(res: FoodItemPortionSizeResponse[]): PortionSize[] {
        return res.map((r) => this.fromResponse(r));
    }
    static async add(ps: PortionSizeForm, foodItemId: number): Promise<Result<FoodItemPortionSizeResponse>> {
        const client = getFoodItemsClient();
        return await tryCatch(
            client.apiFoodItemsIdPortionsPost({
                id: foodItemId,
                foodItemPortionPostRequest: {
                    name: ps.name,
                    amount: ps.amount,
                },
            }),
        );
    }
}

export type PortionSizeForm = {
    amount: number;
    name: string;
};

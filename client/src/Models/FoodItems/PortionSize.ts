import type { PortionSizeResponse } from "../../Gen";
import { getFoodItemsClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";

export class PortionSize {
    id: number;
    amount: number;
    name: string;
    constructor() {
        this.id = 0;
        this.amount = 0;
        this.name = "";
    }
    static fromResponse(res: PortionSizeResponse): PortionSize {
        const p = new PortionSize();
        p.id = res.id;
        p.amount = res.amount;
        p.name = res.name;
        return p;
    }
    static fromResponses(res: PortionSizeResponse[]): PortionSize[] {
        return res.map((r) => this.fromResponse(r));
    }
    static async add(ps: PortionSizeForm, foodItemId: number): Promise<Result<PortionSizeResponse>> {
        const client = getFoodItemsClient();
        return await tryCatch(
            client.apiFoodItemsIdPortionsPost({
                id: foodItemId,
                postFoodItemPortion: {
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

import type { Energy } from "../Common/Energy.ts";

export type FoodItemResponse = {
    id: number;
    manufacturer: string;
    product: string;
    source: string;
    public: boolean;
} & Energy;

import type { Macronutrients } from "./Macronutrients.ts";

export type Energy = {
    kCal: number;
} & Macronutrients;

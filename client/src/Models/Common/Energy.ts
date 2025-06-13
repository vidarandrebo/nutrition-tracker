import type { Macronutrients } from "./Macronutrients.ts";

export type Energy = {
    kCal: number;
} & Macronutrients;

export function SumEnergy(e: Energy[]): Energy {
    const sum: Energy = {
        protein: 0.0,
        carbohydrate: 0.0,
        fat: 0.0,
        kCal: 0.0,
    };
    e.forEach((x) => {
        sum.kCal += x.kCal;
        sum.protein += x.protein;
        sum.carbohydrate += x.carbohydrate;
        sum.fat += x.fat;
    });
    return sum;
}

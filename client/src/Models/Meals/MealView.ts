import type { Energy } from "../Common/Energy.ts";

export type MealView = {
    id: number;
    timestamp: Date;
    entries: MealEntryView[];
}

export type MealEntryView = {
    id: number
    name: string
    amount: number
} & Energy

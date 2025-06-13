import type { Energy } from "../Common/Energy.ts";

export type MealView = {
    id: number;
    timestamp: Date;
    entries: MealEntryView[];
};

export enum EntryType {
    FoodItem,
    Recipe,
}

export type MealEntryView = {
    id: number;
    name: string;
    amount: number;
    entryType: EntryType;
} & Energy;

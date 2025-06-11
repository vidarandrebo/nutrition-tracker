import type { Energy } from "../Common/Energy.ts";

export type RecipeView = {
    id: number;
    name: string;
    entries: RecipeEntryView[];
} & Energy;

export type RecipeEntryView = {
    id: number;
    amount: number;
    name: string;
} & Energy;

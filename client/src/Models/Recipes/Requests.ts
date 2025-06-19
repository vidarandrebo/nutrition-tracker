export type RecipeRequest = {
    name: string;
    entries: RecipeEntryRequest[];
};
export type RecipeEntryRequest = {
    amount: number;
    foodItemId: number;
};

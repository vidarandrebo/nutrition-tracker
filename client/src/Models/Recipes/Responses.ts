export type RecipeResponse = {
    id: number
    name: string;
    entries: RecipeEntryResponse[]
}
export type RecipeEntryResponse = {
    id: number
    amount: number;
    foodItemId: number;
}
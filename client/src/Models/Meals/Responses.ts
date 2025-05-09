export type MealResponse = {
    id: number;
    timestamp: string
    sequenceNumber: number;
    entries: MealEntryResponse[];
}
export type MealEntryResponse = {
    id: number;
    foodItemId: number;
    amount: number;
}
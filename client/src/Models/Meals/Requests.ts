export type PostMealRequest = {
    timestamp: Date;
};

export type PostMealEntryRequest = {
    foodItemId: number;
    recipeId: number;
    amount: number;
};

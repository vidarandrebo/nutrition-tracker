export type PostMealRequest = {
    timeStamp: Date;
};

export type PostMealEntryRequest = {
    foodItemId: number;
    recipeId: number;
    amount: number;
};

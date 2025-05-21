export type MealView = {
    id: number;
    timestamp: Date;
    entries: MealEntryView[];
}

export type MealEntryView = {
    id: number
    name: string
    protein: number
    carbohydrate: number
    fat: number
    kCal: number
}
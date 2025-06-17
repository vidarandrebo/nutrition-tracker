import { EnergyView } from "../Common/EnergyView.ts";

export class MealView {
    id: number;
    timestamp: Date;
    entries: MealEntryView[];

    constructor(id: number, timeStamp: Date, entries: MealEntryView[]) {
        this.id = id;
        this.timestamp = timeStamp;
        this.entries = entries;
    }
};

export enum EntryType {
    FoodItem,
    Recipe,
}

export class MealEntryView extends EnergyView {
    id: number;
    name: string;
    amount: number;
    entryType: EntryType;

    constructor(id: number, name: string, amount: number, entryType: EntryType, protein: number, carbohydrate: number, fat: number, kCal: number) {
        super(protein, carbohydrate, fat, kCal);
        this.id = id;
        this.name = name;
        this.amount = amount;
        this.entryType = entryType;
    }
}

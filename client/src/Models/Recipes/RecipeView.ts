import { EnergyView } from "../Common/EnergyView.ts";

export class RecipeView {
    id: number;
    name: string;
    entries: RecipeEntryView[];
    constructor(id: number, name: string, entries: RecipeEntryView[]) {
        this.id = id;
        this.name = name;
        this.entries = entries;
    }
    get protein(): number {
        let sum = 0.0;
        for (const entry of this.entries) {
            sum += entry.protein;
        }
        return sum;
    }
    get carbohydrate(): number {
        let sum = 0.0;
        for (const entry of this.entries) {
            sum += entry.carbohydrate;
        }
        return sum;
    }
    get fat(): number {
        let sum = 0.0;
        for (const entry of this.entries) {
            sum += entry.fat;
        }
        return sum;
    }
    get kCal(): number {
        let sum = 0.0;
        for (const entry of this.entries) {
            sum += entry.kCal;
        }
        return sum;
    }
    get Protein(): string {
        return this.protein.toFixed(2);
    }
    get Carbohydrate(): string {
        return this.carbohydrate.toFixed(2);
    }
    get Fat(): string {
        return this.fat.toFixed(2);
    }
    get KCal(): string {
        return this.kCal.toFixed(0);
    }
}

export class RecipeEntryView extends EnergyView {
    id: number;
    amount: number;
    name: string;
    constructor(
        id: number,
        name: string,
        amount: number,
        protein: number,
        carbohydrate: number,
        fat: number,
        kCal: number,
    ) {
        super(protein, carbohydrate, fat, kCal);
        this.id = id;
        this.name = name;
        this.amount = amount;
    }
}

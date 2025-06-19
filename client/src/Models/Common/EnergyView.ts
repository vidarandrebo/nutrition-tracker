import { MacronutrientsView } from "./MacronutrientsView.ts";
import type { Energy } from "./Energy.ts";

export class EnergyView extends MacronutrientsView {
    kCal: number;
    constructor(protein: number, carbohydrate: number, fat: number, kCal: number) {
        super(protein, carbohydrate, fat);
        this.kCal = kCal;
    }

    static fromEnergy(e: Energy): EnergyView {
        return new EnergyView(e.protein, e.carbohydrate, e.fat, e.kCal);
    }
    get KCal(): string {
        return this.kCal.toFixed(0);
    }
}

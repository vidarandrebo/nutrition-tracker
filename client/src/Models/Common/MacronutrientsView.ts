import type { Macronutrients } from "./Macronutrients.ts";

export class MacronutrientsView {
    protein: number;
    carbohydrate: number;
    fat: number;

    constructor(protein: number, carbohydrate: number, fat: number) {
        this.protein = protein;
        this.carbohydrate = carbohydrate;
        this.fat = fat;
    }

    static fromMacronutrients(m: Macronutrients): MacronutrientsView {
        return new MacronutrientsView(m.protein, m.carbohydrate, m.fat);
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
}

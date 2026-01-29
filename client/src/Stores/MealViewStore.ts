import { defineStore } from "pinia";
import { useMealStore } from "./MealStore.ts";
import { useFoodItemStore } from "./FoodItemStore.ts";
import { computed } from "vue";
import { EntryType, MealEntryView, MealView } from "../Models/Meals/MealView.ts";
import { useRecipeViewStore } from "./RecipeViewStore.ts";
import { EnergyView } from "../Models/Common/EnergyView.ts";

export const useMealViewStore = defineStore("mealViewStore", () => {
    const mealStore = useMealStore();
    const foodItemStore = useFoodItemStore();
    const recipeStore = useRecipeViewStore();

    async function init() {
        await recipeStore.init();
    }

    const mealsView = computed((): MealView[] => {
        return mealStore.mealsForDay.map((m) => {
            const entries: MealEntryView[] = [];
            const foodItemEntries = m.foodItemEntries
                .map((me) => {
                    const fi = foodItemStore.getFoodItem(me.foodItemId);
                    if (fi) {
                        return new MealEntryView(
                            me.id,
                            fi.name,
                            me.amount,
                            EntryType.FoodItem,
                            (fi.protein * me.amount) / 100,
                            (fi.carbohydrate * me.amount) / 100,
                            (fi.fat * me.amount) / 100,
                            (fi.kCal * me.amount) / 100,
                        );
                    }
                })
                .filter((item): item is MealEntryView => !!item);
            const recipeEntries = m.recipeEntries
                .map((me) => {
                    const r = recipeStore.getRecipe(me.recipeId);
                    if (r) {
                        return new MealEntryView(
                            me.id,
                            r.name,
                            me.amount,
                            EntryType.Recipe,
                            r.protein * me.amount,
                            r.carbohydrate * me.amount,
                            r.fat * me.amount,
                            r.kCal * me.amount,
                        );
                    }
                })
                .filter((item): item is MealEntryView => !!item);
            const macroEntries = m.macronutrientEntries
                .map((me) => {
                    if (me) {
                        return new MealEntryView(
                            me.id,
                            "Macros",
                            1,
                            EntryType.Macronutrient,
                            me.protein,
                            me.carbohydrate,
                            me.fat,
                            me.kCal,
                        );
                    }
                })
                .filter((item): item is MealEntryView => !!item);
            entries.push(...recipeEntries);
            entries.push(...foodItemEntries);
            entries.push(...macroEntries);
            return new MealView(m.id, m.timestamp, entries);
        });
    });

    const dailyMacros = computed((): EnergyView => {
        const energy = {
            protein: 0.0,
            carbohydrate: 0.0,
            fat: 0.0,
            kCal: 0.0,
        };
        for (const meal of mealsView.value) {
            for (const entry of meal.entries) {
                energy.protein += entry.protein;
                energy.carbohydrate += entry.carbohydrate;
                energy.fat += entry.fat;
                energy.kCal += entry.kCal;
            }
        }
        return EnergyView.fromEnergy(energy);
    });

    return { mealsView, dailyMacros, init };
});

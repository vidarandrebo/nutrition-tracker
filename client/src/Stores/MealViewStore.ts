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
            return new MealView(
                m.id,
                m.timestamp,
                m.entries.map((me) => {
                    if (me.foodItemId) {
                        const fi = foodItemStore.getFoodItem(me.foodItemId);
                        return new MealEntryView(
                            me.id,
                            fi?.name ?? "",
                            me.amount,
                            EntryType.Recipe,
                            fi ? (fi.protein * me.amount) / 100 : 0.0,
                            fi ? (fi.carbohydrate * me.amount) / 100 : 0.0,
                            fi ? (fi.fat * me.amount) / 100 : 0.0,
                            fi ? (fi.kCal * me.amount) / 100 : 0.0,
                        );
                    } else if (me.recipeId) {
                        const r = recipeStore.getRecipe(me.recipeId);
                        return new MealEntryView(
                            me.id,
                            r?.name ?? "",
                            me.amount,
                            EntryType.Recipe,
                            r ? r.protein * me.amount : 0.0,
                            r ? r.carbohydrate * me.amount : 0.0,
                            r ? r.fat * me.amount : 0.0,
                            r ? r.kCal * me.amount : 0.0
                        );
                    } else {
                        throw new Error("neither food item or recipe found on meal entry");
                    }
                }));
        });
    });

    const dailyMacros = computed((): EnergyView => {
        const energy = {
            protein: 0.0,
            carbohydrate: 0.0,
            fat: 0.0,
            kCal: 0.0
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

import { defineStore } from "pinia";
import { useMealStore } from "./MealStore.ts";
import { useFoodItemStore } from "./FoodItemStore.ts";
import { computed } from "vue";
import type { MealView } from "../Models/Meals/MealView.ts";
import type { Energy } from "../Models/Common/Energy.ts";

export const useMealViewStore = defineStore("mealViewStore", () => {
    const mealStore = useMealStore();
    const foodItemStore = useFoodItemStore();

    const mealsView = computed((): MealView[] => {
        return mealStore.mealsForDay.map((m) => {
            return {
                id: m.id,
                timestamp: m.timestamp,
                entries: m.entries.map((me) => {
                    const fi = foodItemStore.getFoodItem(me.foodItemId);
                    return {
                        id: me.id,
                        name: fi?.name ?? "",
                        protein: fi ? (fi.protein * me.amount) / 100 : 0.0,
                        carbohydrate: fi ? (fi.carbohydrate * me.amount) / 100 : 0.0,
                        fat: fi ? (fi.fat * me.amount) / 100 : 0.0,
                        kCal: fi ? (fi.kCal * me.amount) / 100 : 0.0,
                        amount: me.amount,
                    };
                }),
            };
        });
    });

    const dailyMacros = computed((): Energy => {
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
        return energy;
    });

    return { mealsView, dailyMacros };
});

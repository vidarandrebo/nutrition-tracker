import { defineStore } from "pinia";
import { useFoodItemStore } from "./FoodItemStore.ts";
import { useRecipeStore } from "./RecipeStore.ts";
import { computed } from "vue";
import type { RecipeEntryView, RecipeView } from "../Models/Recipes/RecipeView.ts";
import { SumEnergy } from "../Models/Common/Energy.ts";

export const useRecipeViewStore = defineStore("recipeViewStore", () => {
    const recipeStore = useRecipeStore();
    const foodItemStore = useFoodItemStore();
    const recipesView = computed(() => {
        return recipeStore.collection.map((r): RecipeView => {
            const entries = r.entries.map((e): RecipeEntryView => {
                const fi = foodItemStore.getFoodItem(e.foodItemId);
                return {
                    amount: e.amount,
                    protein: fi ? (fi.protein * e.amount) / 100.0 : 0.0,
                    carbohydrate: fi ? (fi.carbohydrate * e.amount) / 100.0 : 0.0,
                    fat: fi ? (fi.fat * e.amount) / 100.0 : 0.0,
                    kCal: fi ? (fi.kCal * e.amount) / 100.0 : 0.0,
                    name: fi ? fi.name : "",
                    id: e.id,
                };
            });
            const e = SumEnergy(entries);
            return {
                id: r.id,
                name: r.name,
                entries: entries,
                ...e,
            };
        });
    });

    function getRecipe(id: number): RecipeView | undefined {
        const recipe = recipeStore.collection.find((r) => r.id === id);
        if (!recipe) {
            return undefined;
        }
        const entries = recipe.entries.map((e): RecipeEntryView => {
            const fi = foodItemStore.getFoodItem(e.foodItemId);
            return {
                amount: e.amount,
                protein: fi ? (fi.protein * e.amount) / 100.0 : 0.0,
                carbohydrate: fi ? (fi.carbohydrate * e.amount) / 100.0 : 0.0,
                fat: fi ? (fi.fat * e.amount) / 100.0 : 0.0,
                kCal: fi ? (fi.kCal * e.amount) / 100.0 : 0.0,
                name: fi ? fi.name : "",
                id: e.id,
            };
        });
        const e = SumEnergy(entries);
        return {
            id: recipe.id,
            name: recipe.name,
            entries: entries,
            ...e,
        };
    }

    async function init() {
        await Promise.all([recipeStore.init(), foodItemStore.init()]);
    }

    return { recipesView, init, getRecipe };
});

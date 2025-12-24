import { defineStore } from "pinia";
import { useFoodItemStore } from "./FoodItemStore.ts";
import { useRecipeStore } from "./RecipeStore.ts";
import { computed } from "vue";
import { RecipeEntryView, RecipeView } from "../Models/Recipes/RecipeView.ts";

export const useRecipeViewStore = defineStore("recipeViewStore", () => {
    const recipeStore = useRecipeStore();
    const foodItemStore = useFoodItemStore();
    const recipesView = computed(() => {
        return recipeStore.collection.map((r): RecipeView => {
            const entries = r.foodItemEntries.map((e): RecipeEntryView => {
                const fi = foodItemStore.getFoodItem(e.foodItemId);
                return new RecipeEntryView(
                    e.id,
                    fi ? fi.name : "",
                    e.amount,
                    fi ? (fi.protein * e.amount) / 100.0 : 0.0,
                    fi ? (fi.carbohydrate * e.amount) / 100.0 : 0.0,
                    fi ? (fi.fat * e.amount) / 100.0 : 0.0,
                    fi ? (fi.kCal * e.amount) / 100.0 : 0.0,
                );
            });
            return new RecipeView(r.id, r.name, entries);
        });
    });

    function getRecipe(id: number): RecipeView | undefined {
        const recipe = recipeStore.collection.find((r) => r.id === id);
        if (!recipe) {
            return undefined;
        }
        const entries = recipe.foodItemEntries.map((e): RecipeEntryView => {
            const fi = foodItemStore.getFoodItem(e.foodItemId);
            return new RecipeEntryView(
                e.id,
                fi ? fi.name : "",
                e.amount,
                fi ? (fi.protein * e.amount) / 100.0 : 0.0,
                fi ? (fi.carbohydrate * e.amount) / 100.0 : 0.0,
                fi ? (fi.fat * e.amount) / 100.0 : 0.0,
                fi ? (fi.kCal * e.amount) / 100.0 : 0.0,
            );
        });
        return new RecipeView(recipe.id, recipe.name, entries);
    }

    async function init() {
        await Promise.all([recipeStore.init(), foodItemStore.init()]);
    }

    return { recipesView, init, getRecipe };
});

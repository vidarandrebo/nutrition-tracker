import { defineStore } from "pinia";
import { Recipe } from "../Models/Recipes/Recipe.ts";
import { ref } from "vue";
import type { RecipeRequest } from "../Models/Recipes/Requests.ts";

export const useRecipeStore = defineStore("recipeStore", () => {
    const collection = ref<Recipe[]>([]);
    const initialized = ref<boolean>(false);

    async function init() {
        if (!initialized.value) {
            const items = await Recipe.get();
            if (items === null) {
                collection.value = [];
            }
            else {
                collection.value = items;
            }
            initialized.value = true
        }

    }

    async function addRecipe(recipe: RecipeRequest) {
        const newRecipe = await Recipe.add(recipe);
        if (newRecipe) {
            collection.value.push(newRecipe)
        }
    }

    return { collection, addRecipe , init};
});
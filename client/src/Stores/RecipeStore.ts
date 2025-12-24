import { defineStore } from "pinia";
import { Recipe } from "../Models/Recipes/Recipe.ts";
import { ref } from "vue";
import { Failure, type Result, Success } from "../Utilities/tryCatch.ts";
import type { RecipePostRequest } from "../Gen";

export const useRecipeStore = defineStore("recipeStore", () => {
    const collection = ref<Recipe[]>([]);
    const initialized = ref<boolean>(false);

    async function init() {
        if (!initialized.value) {
            const items = await Recipe.get();
            if (items === null) {
                collection.value = [];
            } else {
                collection.value = items;
            }
            initialized.value = true;
        }
    }

    async function addRecipe(recipe: RecipePostRequest) {
        const newRecipe = await Recipe.add(recipe);
        if (newRecipe) {
            collection.value.push(newRecipe);
        }
    }
    function removeRecipe(id: number): Result<void> {
        const idx = collection.value.findIndex((m) => m.id === id);
        if (idx !== -1) {
            collection.value.splice(idx, 1);
            return Success.empty();
        }
        return Failure.new(new Error("failed to remove recipe"));
    }

    return { collection, addRecipe, init, removeRecipe };
});

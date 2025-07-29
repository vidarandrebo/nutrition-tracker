<script setup lang="ts">
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import router from "../../Router.ts";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { onMounted } from "vue";
import { useRecipeViewStore } from "../../Stores/RecipeViewStore.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import RecipeView from "./RecipeView.vue";
import { Recipe } from "../../Models/Recipes/Recipe.ts";

const recipeStore = useRecipeStore();
const foodItemStore = useFoodItemStore();

const recipeViewStore = useRecipeViewStore();
onMounted(async () => {
    await Promise.all([recipeStore.init(), foodItemStore.init()]);
});

async function addRecipe() {
    await router.push("/recipes/add");
}
async function onDeleteRecipe(recipeId: number) {
    const { error } = await Recipe.delete(recipeId);
    if (error) {
        return;
    }

    const { error: rmRecipeErr } = recipeStore.removeRecipe(recipeId);
    if (rmRecipeErr) {
        console.log(rmRecipeErr.message);
    }
}
</script>

<template>
    <div class="container">
        <LevelPrimary>
            <template #left>
                <HeaderH1 class="level-item">Recipes</HeaderH1>
            </template>
            <template #right>
                <ButtonPrimary class="level-item" @click="addRecipe">Add</ButtonPrimary>
            </template>
        </LevelPrimary>
        <article v-for="item in recipeViewStore.recipesView" :key="item.id" class="box">
            <RecipeView :item="item" @delete-recipe="onDeleteRecipe"></RecipeView>
        </article>
    </div>
</template>

<style scoped></style>

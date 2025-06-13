<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import router from "../../Router.ts";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { onMounted } from "vue";
import { useRecipeViewStore } from "../../Stores/RecipeViewStore.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";

const recipeStore = useRecipeStore();
const foodItemStore = useFoodItemStore();

const recipeViewStore = useRecipeViewStore();
onMounted(async () => {
    await Promise.all([recipeStore.init(), foodItemStore.init()]);
});

async function addRecipe() {
    await router.push("/recipes/add");
}
</script>

<template>
    <LevelPrimary>
        <template #left>
            <HeaderH1 class="level-item">Recipes</HeaderH1>
        </template>
        <template #right>
            <ButtonPrimary class="level-item" @click="addRecipe">Add</ButtonPrimary>
        </template>
    </LevelPrimary>
    <article v-for="item in recipeViewStore.recipesView" :key="item.id" class="box">
        {{ item.name }}
        <p>
            KCal: {{ item.kCal }}, Protein: {{ item.protein }} g, Carbohydrate: {{ item.carbohydrate }} g, Fat:
            {{ item.fat }} g
        </p>
    </article>
</template>

<style scoped></style>

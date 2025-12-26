<script setup lang="ts">
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useRoute } from "vue-router";
import { useMealStore } from "../../Stores/MealStore.ts";
import { onMounted, ref } from "vue";
import type { Meal } from "../../Models/Meals/Meal.ts";
import TabMenu from "../../Components/TabMenu.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import FoodItemTab from "./FoodItemTab.vue";
import RecipeTab from "./RecipeTab.vue";
import type { MealFoodItemEntryPostRequest, MealRecipeEntryPostRequest } from "../../Gen";

const activeTab = ref<string>("Food Items");
const mealStore = useMealStore();
const foodItemStore = useFoodItemStore();

async function addRecipeToMeal(form: MealRecipeEntryPostRequest) {
    if (meal.value) {
        await mealStore.addRecipeEntry(form, meal.value.id);
    }
}
async function addFoodItemToMeal(form: MealFoodItemEntryPostRequest) {
    if (meal.value) {
        await mealStore.addFoodItemEntry(form, meal.value.id);
    }
}

const route = useRoute();
let mealId = 0;

if (route.params.id && !Array.isArray(route.params.id)) {
    mealId = parseInt(route.params.id);
}

const meal = ref<Meal | null>(null);

onMounted(async () => {
    let m = mealStore.getMeal(mealId);
    if (!m) {
        await mealStore.loadMeal(mealId);
        m = mealStore.getMeal(mealId);
    }
    meal.value = m;
    await foodItemStore.init();
});
</script>

<template>
    <section class="container">
        <HeaderH1>Meal {{ mealId }}</HeaderH1>
        <div v-if="meal">
            <p>{{ meal.timestamp }}</p>
        </div>
        <div v-else>
            <p class="is-warning">No meal found</p>
        </div>
        <TabMenu
            :entries="['Food Items', 'Recipes']"
            preselected="Food Items"
            @selected="(value) => (activeTab = value)"
        ></TabMenu>
        <template v-if="activeTab === 'Food Items'">
            <FoodItemTab @add-entry="addFoodItemToMeal"></FoodItemTab>
        </template>
        <template v-if="activeTab === 'Recipes'">
            <RecipeTab @add-entry="addRecipeToMeal"></RecipeTab>
        </template>
    </section>
</template>

<style scoped></style>

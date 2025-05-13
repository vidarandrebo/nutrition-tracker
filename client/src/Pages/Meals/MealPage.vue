<script setup lang="ts">

import HeaderH1 from "../../Components/HeaderH1.vue";
import { useRoute } from "vue-router";
import { useMealStore } from "../../Stores/MealStore.ts";
import { onMounted, ref, watch } from "vue";
import type { Meal } from "../../Models/Meals/Meal.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import debounce from "debounce";
import InputText from "../../Components/InputText.vue";
import FormField from "../../Components/FormField.vue";
import FoodItemDisplay from "../FoodItems/FoodItemDisplay.vue";

const mealStore = useMealStore();

const foodItemStore = useFoodItemStore();

const route = useRoute();
let mealId = 0;

const searchTerm = ref<string>("");

if (!Array.isArray(route.params.id)) {
    mealId = parseInt(route.params.id)
}

const meal = ref<Meal | null>(null)


const updateSearchTermDb = debounce(() => {
    foodItemStore.searchTerm = searchTerm.value
}, 400);

watch(searchTerm, () => {
    updateSearchTermDb();
});

onMounted(async () => {
    meal.value = mealStore.getMeal(mealId)
    await foodItemStore.init()
})



</script>

<template>
<HeaderH1>Meal {{ mealId }}</HeaderH1>
    <div v-if="meal">
        <p>{{ meal.timestamp }}</p>
    </div>
    <div v-else>
        <p class="is-warning">No meal found</p>
    </div>
    <FormField>
        <InputText v-model="searchTerm" placeholder="Search"></InputText>
    </FormField>
    <ul>
        <div v-for="foodItem in foodItemStore.filteredFoodItems" :key="foodItem.id">{{ foodItem.name}}</div>
    </ul>
</template>

<style scoped>

</style>
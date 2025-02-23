<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted } from "vue";
import FoodItemDisplay from "./FoodItemDisplay.vue";

const foodItemStore = useFoodItemStore();

onMounted(async () => {
    await foodItemStore.init();
});
</script>

<template>
    <HeaderH1>Food Items</HeaderH1>
    <RouterLink to="/food-items/add">Add</RouterLink>
    <ul>
        <FoodItemDisplay v-for="foodItem in foodItemStore.collection" :item="foodItem"></FoodItemDisplay>
    </ul>
</template>

<style scoped></style>

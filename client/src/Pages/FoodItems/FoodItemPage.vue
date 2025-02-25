<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { computed, onMounted, ref } from "vue";
import FoodItemDisplay from "./FoodItemDisplay.vue";
import InputText from "../../Components/InputText.vue";

const foodItemStore = useFoodItemStore();

onMounted(async () => {
    await foodItemStore.init();
});
const searchTerm = ref<string>("");
const filteredFoodItems = computed(() => {
    const terms = searchTerm.value
        .split(" ")
        .filter((s) => s !== "")
        .map((t) => t.toLowerCase());

    if (searchTerm.value === "") {
        return foodItemStore.collection;
    }
    return foodItemStore.collection.filter((x) => {
        for (let i = 0; i < terms.length; i++) {
            if (!(x.product.toLowerCase().includes(terms[i]) || x.manufacturer.toLowerCase().includes(terms[i]))) {
                return false;
            }
        }
        return true;
    });
});
</script>

<template>
    <HeaderH1>Food Items</HeaderH1>
    <RouterLink to="/food-items/add">Add</RouterLink>
    <InputText v-model="searchTerm"></InputText>
    <ul>
        <FoodItemDisplay v-for="foodItem in filteredFoodItems" :item="foodItem" :key="foodItem.id"></FoodItemDisplay>
    </ul>
</template>

<style scoped></style>

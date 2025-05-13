<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted, ref, watch } from "vue";
import FoodItemDisplay from "./FoodItemDisplay.vue";
import InputText from "../../Components/InputText.vue";
import FormField from "../../Components/FormField.vue";
import  debounce  from "debounce";

const foodItemStore = useFoodItemStore();
const searchTerm = ref<string>("");

onMounted(async () => {
    await foodItemStore.init();
});

const updateSearchTermDb = debounce(() => {
    foodItemStore.searchTerm = searchTerm.value
}, 400);

watch(searchTerm, () => {
    updateSearchTermDb();
});
</script>

<template>
    <HeaderH1>Food Items</HeaderH1>
    <FormField>
        <RouterLink class="button is-primary" to="/food-items/add">Add</RouterLink>
    </FormField>
    <FormField>
        <InputText v-model="searchTerm" placeholder="Search"></InputText>
    </FormField>
    <ul>
        <FoodItemDisplay
            v-for="foodItem in foodItemStore.filteredFoodItems"
            :item="foodItem"
            :key="foodItem.id"
        ></FoodItemDisplay>
    </ul>
</template>

<style scoped></style>

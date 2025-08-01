<script setup lang="ts">
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted, ref, watch } from "vue";
import FoodItemDisplay from "./FoodItemDisplay.vue";
import InputText from "../../Components/Forms/InputText.vue";
import FormField from "../../Components/Forms/FormField.vue";
import debounce from "debounce";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";

const foodItemStore = useFoodItemStore();
const searchTerm = ref<string>("");

onMounted(async () => {
    await foodItemStore.init();
});

const updateSearchTermDb = debounce(() => {
    foodItemStore.searchTerm = searchTerm.value;
}, 400);

watch(searchTerm, () => {
    updateSearchTermDb();
});
async function onDeleteFoodItem(foodItemId: number) {
    const { error } = await FoodItem.delete(foodItemId);
    if (error) {
        return;
    }

    const { error: rmFoodItemErr } = foodItemStore.removeFoodItem(foodItemId);
    if (rmFoodItemErr) {
        console.log(rmFoodItemErr.message);
    }
}
</script>

<template>
    <section class="container">
        <LevelPrimary>
            <template #left>
                <HeaderH1>Food Items</HeaderH1>
            </template>
            <template #right>
                <FormField>
                    <RouterLink class="button is-primary" to="/food-items/add">Add</RouterLink>
                </FormField>
            </template>
        </LevelPrimary>
        <FormField>
            <InputText v-model="searchTerm" placeholder="Search"></InputText>
        </FormField>
        <ul>
            <FoodItemDisplay
                v-for="foodItem in foodItemStore.filteredFoodItems"
                :key="foodItem.id"
                :item="foodItem"
                @delete-food-item="onDeleteFoodItem"
            ></FoodItemDisplay>
        </ul>
    </section>
</template>

<style scoped></style>

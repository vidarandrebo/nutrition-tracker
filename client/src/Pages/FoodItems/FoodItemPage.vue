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
import { useFilterStore } from "../../Stores/FilterStore.ts";
import FoodItemTab from "../Meals/FoodItemTab.vue";
import FoodItemTable from "./FoodItemTable.vue";

const foodItemStore = useFoodItemStore();
const filterStore = useFilterStore();
const searchTerm = ref<string>(filterStore.foodItem.searchTerm);

onMounted(async () => {
    await foodItemStore.init();
});

const updateSearchTermDb = debounce(() => {
    filterStore.foodItem.searchTerm = searchTerm.value;
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
                    <RouterLink class="button is-primary is-soft" to="/food-items/add">Add</RouterLink>
                </FormField>
            </template>
        </LevelPrimary>
        <div class="is-flex is-align-items-center is-gap-1">
            <FormField class="is-flex-grow-1">
                <InputText v-model="searchTerm" placeholder="Search"></InputText>
            </FormField>
            <FormField>
                <label class="checkbox">
                    Show public
                    <input v-model="filterStore.foodItem.showPublic" type="checkbox" />
                </label>
            </FormField>
        </div>
        <FoodItemTable :food-items="foodItemStore.filteredFoodItems"></FoodItemTable>
        <!--<ul>
            <FoodItemDisplay
                v-for="foodItem in foodItemStore.filteredFoodItems"
                :key="foodItem.id"
                :item="foodItem"
                @delete-food-item="onDeleteFoodItem"
            ></FoodItemDisplay>
        </ul>-->
    </section>
</template>

<style scoped></style>

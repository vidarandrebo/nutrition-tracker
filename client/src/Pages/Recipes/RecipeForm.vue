<script setup lang="ts">
import { ref } from "vue";
import type { RecipeRequest } from "../../Models/Recipes/Requests.ts";
import InputText from "../../Components/Forms/InputText.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import HeaderH2 from "../../Components/HeaderH2.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";
import router from "../../Router.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";

const item = ref<RecipeRequest>({ name: "", entries: [] });
const showFoodItemSelector = ref<boolean>(false);
const recipeStore = useRecipeStore();
const foodItemStore = useFoodItemStore();

function addEntry() {
    showFoodItemSelector.value = true;
}

function onFoodItemSelected(id: number) {
    item.value.entries.push({ amount: 100, foodItemId: id });
    showFoodItemSelector.value = false;
}

async function submit() {
    await recipeStore.addRecipe(item.value);
    await router.push("/recipes");
}
</script>

<template>
    <form @submit.prevent="submit">
        <LabelPrimary>
            Name
            <InputText v-model="item.name"></InputText>
        </LabelPrimary>
        <LevelPrimary v-if="!showFoodItemSelector">
            <template #left>
                <HeaderH2 class="level-item">Entries</HeaderH2>
            </template>
            <template #right>
                <ButtonPrimary @click="addEntry">Add Entry</ButtonPrimary>
            </template>
        </LevelPrimary>
        <FoodItemSelector
            v-if="showFoodItemSelector"
            @select="onFoodItemSelected"
            @cancel="showFoodItemSelector = false"
        ></FoodItemSelector>
        <template v-for="(_, id) in item.entries" v-else :key="id">
            <div class="box is-flex is-flex-direction-row is-justify-content-space-between">
                <LabelPrimary>
                    FoodItem
                    <InputNumber v-model="item.entries[id].foodItemId" disabled></InputNumber>
                    {{ foodItemStore.getFoodItem(item.entries[id].foodItemId)?.name }}
                </LabelPrimary>
                <LabelPrimary>
                    Amount
                    <InputNumber v-model="item.entries[id].amount"></InputNumber>
                </LabelPrimary>
            </div>
        </template>
        <ButtonPrimary type="submit">Save</ButtonPrimary>
    </form>
    <p>{{ item }}</p>
</template>

<style scoped></style>

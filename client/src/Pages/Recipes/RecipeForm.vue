<script setup lang="ts">
import { ref } from "vue";
import type { RecipeRequest } from "../../Models/Recipes/Requests.ts";
import InputText from "../../Components/Forms/InputText.vue";
import Label from "../../Components/Forms/Label.vue";
import HeaderH2 from "../../Components/HeaderH2.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";
import Level from "../../Components/Level.vue";
import { Recipe } from "../../Models/Recipes/Recipe.ts";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";

const item = ref<RecipeRequest>({ name: "", entries: [] });
const showFoodItemSelector = ref<boolean>(false);
const recipeStore = useRecipeStore();

function addEntry() {
    showFoodItemSelector.value = true;
}

function onFoodItemSelected(id: number) {
    item.value.entries.push({ amount: 100, foodItemId: id });
    showFoodItemSelector.value = false;
}

async function submit() {
    await recipeStore.addRecipe(item.value)
}
</script>

<template>
    <form v-on:submit.prevent="submit">
        <Label>
            Name
            <InputText v-model="item.name"></InputText>
        </Label>
        <Level v-if="!showFoodItemSelector">
            <template #left>
                <HeaderH2 class="level-item">Entries</HeaderH2>
            </template>
            <template #right>
                <ButtonPrimary @click="addEntry">Add Entry</ButtonPrimary>
            </template>
        </Level>
        <FoodItemSelector v-if="showFoodItemSelector" @select="onFoodItemSelected" @cancel="showFoodItemSelector = false"></FoodItemSelector>
        <template v-for="(_, id) in item.entries" v-else>
            <div class="box is-flex is-flex-direction-row is-justify-content-space-between">
                <Label>
                    FoodItem
                    <InputNumber v-model="item.entries[id].foodItemId" disabled></InputNumber>
                </Label>
                <Label>
                    Amount
                    <InputNumber v-model="item.entries[id].amount"></InputNumber>
                </Label>
            </div>
        </template>
        <ButtonPrimary type="submit">Save</ButtonPrimary>
    </form>
    <p>{{ item }}</p>
</template>

<style scoped>

</style>
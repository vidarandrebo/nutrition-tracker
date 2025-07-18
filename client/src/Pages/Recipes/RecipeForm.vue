<script setup lang="ts">
import { computed, ref } from "vue";
import type { RecipeRequest } from "../../Models/Recipes/Requests.ts";
import InputText from "../../Components/Forms/InputText.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import HeaderH2 from "../../Components/Headings/HeaderH2.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";
import router from "../../Router.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import FormField from "../../Components/Forms/FormField.vue";

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

const saveEnabled = computed((): boolean => {
    return item.value.name !== "" && item.value.entries.length > 0;
});
const saveHelpText = computed((): string => {
    return "Recipe needs to have a name and have at least one foodItem";
});
</script>

<template>
    <form @submit.prevent="submit">
        <FormField>
            <LabelPrimary>
                Name
                <InputText v-model="item.name"></InputText>
            </LabelPrimary>
        </FormField>
        <LevelPrimary v-if="!showFoodItemSelector">
            <template #left>
                <HeaderH2 class="level-item">Ingredients</HeaderH2>
            </template>
            <template #right>
                <ButtonPrimary @click="addEntry">Add Ingredient</ButtonPrimary>
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
        <LevelPrimary>
            <template #right>
                <ButtonPrimary type="submit" :enabled="saveEnabled" :disabled-text="saveHelpText">Save</ButtonPrimary>
            </template>
        </LevelPrimary>
    </form>
</template>

<style scoped></style>

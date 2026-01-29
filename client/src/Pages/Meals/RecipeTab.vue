<script setup lang="ts">
import InputNumber from "../../Components/Forms/InputNumber.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import ModalPrimary from "../../Components/ModalPrimary.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import { computed, ref } from "vue";
import type { Energy } from "../../Models/Common/Energy.ts";
import RecipeSelector from "../Recipes/RecipeSelector.vue";
import { useRecipeViewStore } from "../../Stores/RecipeViewStore.ts";
import type { RecipeView } from "../../Models/Recipes/RecipeView.ts";
import type { MealRecipeEntryPostRequest } from "../../Gen";

const recipeStore = useRecipeViewStore();
const selectedRecipe = ref<RecipeView | undefined>(undefined);
const modalActive = ref<boolean>(false);
const recipeForm = ref<MealRecipeEntryPostRequest>({
    recipeId: 0,
    amount: 0,
});

const emit = defineEmits<{
    addEntry: [entry: MealRecipeEntryPostRequest];
}>();

function showItemDialog(itemId: number) {
    recipeForm.value.amount = 1.0;
    recipeForm.value.recipeId = itemId;
    selectedRecipe.value = recipeStore.getRecipe(itemId);
    modalActive.value = true;
}

const nutrients = computed((): Energy => {
    if (selectedRecipe.value) {
        return {
            kCal: selectedRecipe.value.kCal * recipeForm.value.amount,
            protein: selectedRecipe.value.protein * recipeForm.value.amount,
            carbohydrate: selectedRecipe.value.carbohydrate * recipeForm.value.amount,
            fat: selectedRecipe.value.fat * recipeForm.value.amount,
        };
    }
    return { kCal: 0, protein: 0, carbohydrate: 0, fat: 0 };
});

function submit() {
    emit("addEntry", recipeForm.value);
    modalActive.value = false;
}
</script>

<template>
    <RecipeSelector @select="(item) => showItemDialog(item)"></RecipeSelector>
    <ModalPrimary v-model="modalActive" :title="selectedRecipe?.name">
        <template #default>
            <div class="is-flex is-flex-direction-column">
                <LabelPrimary>
                    <p>Amount (g)</p>
                    <InputNumber v-model="recipeForm.amount"></InputNumber>
                </LabelPrimary>
                <div class="is-flex is-flex-direction-row is-justify-content-space-between is-flex-wrap-wrap">
                    <p class="pr-2"><b>KCal:</b> {{ nutrients.kCal }}</p>
                    <p class="pr-2"><b>Protein:</b> {{ nutrients.protein }}&nbsp;g</p>
                    <p class="pr-2"><b>Carbohydrate:</b> {{ nutrients.carbohydrate }}&nbsp;g</p>
                    <p class="pr-2"><b>Fat:</b> {{ nutrients.fat }}&nbsp;g</p>
                </div>
            </div>
        </template>
        <template #footer>
            <ButtonPrimary @click="submit">Add</ButtonPrimary>
        </template>
    </ModalPrimary>
</template>

<style scoped></style>

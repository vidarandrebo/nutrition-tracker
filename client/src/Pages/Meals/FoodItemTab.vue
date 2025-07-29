<script setup lang="ts">
import InputNumber from "../../Components/Forms/InputNumber.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import ModalPrimary from "../../Components/ModalPrimary.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import { computed, ref, watch } from "vue";
import type { Energy } from "../../Models/Common/Energy.ts";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import debounce from "debounce";
import type { PostMealEntryRequest } from "../../Models/Meals/Requests.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";

const foodItemStore = useFoodItemStore();
const searchTerm = ref<string>("");
const selectedFoodItem = ref<FoodItem | undefined>(undefined);
const modalActive = ref<boolean>(false);
const foodItemForm = ref<PostMealEntryRequest>({
    foodItemId: 0,
    recipeId: 0,
    amount: 0,
});

const emit = defineEmits<{
    addEntry: [entry: PostMealEntryRequest];
}>();

watch(searchTerm, () => {
    updateSearchTermDb();
});
const updateSearchTermDb = debounce(() => {
    foodItemStore.searchTerm = searchTerm.value;
}, 400);

function showItemDialog(itemId: number) {
    foodItemForm.value.amount = 100;
    foodItemForm.value.foodItemId = itemId;
    selectedFoodItem.value = foodItemStore.getFoodItem(itemId);
    modalActive.value = true;
}

const nutrients = computed((): Energy => {
    if (selectedFoodItem.value) {
        return {
            kCal: (selectedFoodItem.value.kCal * foodItemForm.value.amount) / 100,
            protein: (selectedFoodItem.value.protein * foodItemForm.value.amount) / 100,
            carbohydrate: (selectedFoodItem.value.carbohydrate * foodItemForm.value.amount) / 100,
            fat: (selectedFoodItem.value.fat * foodItemForm.value.amount) / 100,
        };
    }
    return { kCal: 0, protein: 0, carbohydrate: 0, fat: 0 };
});

function submit() {
    emit("addEntry", foodItemForm.value);
    modalActive.value = false;
}
</script>

<template>
    <FoodItemSelector @select="(item) => showItemDialog(item)"></FoodItemSelector>
    <ModalPrimary v-model="modalActive" :title="selectedFoodItem?.name">
        <template #default>
            <div class="is-flex is-flex-direction-column">
                <LabelPrimary>
                    <p>Amount (g)</p>
                    <InputNumber v-model="foodItemForm.amount"></InputNumber>
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

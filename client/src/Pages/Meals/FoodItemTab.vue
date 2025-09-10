<script setup lang="ts">
import InputNumber from "../../Components/Forms/InputNumber.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import ModalPrimary from "../../Components/ModalPrimary.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import { computed, ref, watch } from "vue";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import type { PostMealEntryRequest } from "../../Models/Meals/Requests.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { EnergyView } from "../../Models/Common/EnergyView.ts";

const foodItemStore = useFoodItemStore();
const amount = ref<number>(1.0);
const unitMultiplier = ref<number>(100);
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

function showItemDialog(itemId: number) {
    amount.value = 1;
    updateFoodItemAmount();
    foodItemForm.value.foodItemId = itemId;
    selectedFoodItem.value = foodItemStore.getFoodItem(itemId);
    modalActive.value = true;
}

const nutrients = computed((): EnergyView => {
    if (selectedFoodItem.value) {
        return EnergyView.fromEnergy({
            kCal: (selectedFoodItem.value.kCal * foodItemForm.value.amount) / 100,
            protein: (selectedFoodItem.value.protein * foodItemForm.value.amount) / 100,
            carbohydrate: (selectedFoodItem.value.carbohydrate * foodItemForm.value.amount) / 100,
            fat: (selectedFoodItem.value.fat * foodItemForm.value.amount) / 100,
        });
    }
    return new EnergyView(0, 0, 0, 0);
});
function updateFoodItemAmount() {
    foodItemForm.value.amount = amount.value * unitMultiplier.value;
}

watch(amount, () => {
    updateFoodItemAmount();
});
watch(unitMultiplier, () => {
    updateFoodItemAmount();
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
                <div class="grid">
                    <LabelPrimary>
                        <p>Amount</p>
                        <InputNumber v-model="amount"></InputNumber>
                    </LabelPrimary>
                    <LabelPrimary>
                        <p>Unit</p>
                        <select v-if="selectedFoodItem" v-model="unitMultiplier" class="select">
                            <option selected :value="100">100 Grams</option>
                            <option :value="1">Grams</option>
                            <option v-for="p in selectedFoodItem.portionSizes" :key="p.id" :value="p.amount">
                                {{ p.name }} ({{ p.amount }}g)
                            </option>
                        </select>
                    </LabelPrimary>
                </div>
                <div class="is-flex is-flex-direction-row is-justify-content-space-between is-flex-wrap-wrap">
                    <p class="pr-2"><b>KCal:</b> {{ nutrients.KCal }}</p>
                    <p class="pr-2"><b>Protein:</b> {{ nutrients.Protein }}&nbsp;g</p>
                    <p class="pr-2"><b>Carbohydrate:</b> {{ nutrients.Carbohydrate }}&nbsp;g</p>
                    <p class="pr-2"><b>Fat:</b> {{ nutrients.Fat }}&nbsp;g</p>
                </div>
            </div>
        </template>
        <template #footer>
            <ButtonPrimary @click="submit">Add</ButtonPrimary>
        </template>
    </ModalPrimary>
</template>

<style scoped></style>

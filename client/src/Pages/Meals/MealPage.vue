<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useRoute } from "vue-router";
import { useMealStore } from "../../Stores/MealStore.ts";
import { computed, onMounted, ref, watch } from "vue";
import type { Meal } from "../../Models/Meals/Meal.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import debounce from "debounce";
import { FoodItem } from "../../Models/FoodItems/Fooditem.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import type { PostMealEntryRequest } from "../../Models/Meals/Requests.ts";
import Label from "../../Components/Forms/Label.vue";
import Modal from "../../Components/Modal.vue";
import type { Energy } from "../../Models/Common/Energy.ts";
import TabMenu from "../../Components/TabMenu.vue";
import FoodItemSelector from "../FoodItems/FoodItemSelector.vue";

const mealStore = useMealStore();

const activeTab = ref<string>("Food Items");

const foodItemForm = ref<PostMealEntryRequest>({
    foodItemId: 0,
    amount: 0,
});

const foodItemStore = useFoodItemStore();

const route = useRoute();
let mealId = 0;

const searchTerm = ref<string>("");

if (!Array.isArray(route.params.id)) {
    mealId = parseInt(route.params.id);
}

const meal = ref<Meal | null>(null);

const selectFoodItem = ref(true);

const selectedFoodItem = ref<FoodItem | undefined>(undefined);

const updateSearchTermDb = debounce(() => {
    foodItemStore.searchTerm = searchTerm.value;
}, 400);

watch(searchTerm, () => {
    updateSearchTermDb();
});

onMounted(async () => {
    let m = mealStore.getMeal(mealId);
    if (!m) {
        await mealStore.loadMeal(mealId);
        m = mealStore.getMeal(mealId);
    }
    meal.value = m;
    await foodItemStore.init();
});

function showFoodItemDialog(itemId: number) {
    selectFoodItem.value = false;
    foodItemForm.value.amount = 100;
    foodItemForm.value.foodItemId = itemId;
    selectedFoodItem.value = foodItemStore.getFoodItem(itemId);
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

async function addToMeal() {
    if (meal.value) {
        await mealStore.addMealEntry(foodItemForm.value, meal.value.id);
    }
    selectFoodItem.value = true;
}
</script>

<template>
    <HeaderH1>Meal {{ mealId }}</HeaderH1>
    <div v-if="meal">
        <p>{{ meal.timestamp }}</p>
    </div>
    <div v-else>
        <p class="is-warning">No meal found</p>
    </div>
    <TabMenu
        :entries="['Food Items', 'Recipes']"
        preselected="Food Items"
        @selected="(value) => (activeTab = value)"
    ></TabMenu>
    <template v-if="activeTab === 'Food Items'">
        <template v-if="selectFoodItem">
            <FoodItemSelector @select="(item) => showFoodItemDialog(item)"></FoodItemSelector>
        </template>
        <template v-else>
            <Modal @closed="() => (selectFoodItem = true)" :title="selectedFoodItem?.name">
                <template #default>
                    <div class="is-flex is-flex-direction-column">
                        <Label>
                            <p>Amount (g)</p>
                            <InputNumber v-model="foodItemForm.amount"></InputNumber>
                        </Label>
                        <div class="is-flex is-flex-direction-row is-justify-content-space-between is-flex-wrap-wrap">
                            <p class="pr-2"><b>KCal:</b> {{ nutrients.kCal }}</p>
                            <p class="pr-2"><b>Protein:</b> {{ nutrients.protein }}&nbsp;g</p>
                            <p class="pr-2"><b>Carbohydrate:</b> {{ nutrients.carbohydrate }}&nbsp;g</p>
                            <p class="pr-2"><b>Fat:</b> {{ nutrients.fat }}&nbsp;g</p>
                        </div>
                    </div>
                </template>
                <template #footer>
                    <ButtonPrimary @click="addToMeal">Add</ButtonPrimary>
                </template>
            </Modal>
        </template>
    </template>
</template>

<style scoped></style>

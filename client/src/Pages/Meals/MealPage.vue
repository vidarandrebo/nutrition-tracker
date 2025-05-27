<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useRoute } from "vue-router";
import { useMealStore } from "../../Stores/MealStore.ts";
import { computed, onMounted, ref, watch } from "vue";
import type { Meal } from "../../Models/Meals/Meal.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import debounce from "debounce";
import InputText from "../../Components/InputText.vue";
import FormField from "../../Components/FormField.vue";
import { FoodItem } from "../../Models/FoodItems/Fooditem.ts";
import Button from "../../Components/Button.vue";
import InputNumber from "../../Components/InputNumber.vue";
import type { PostMealEntryRequest } from "../../Models/Meals/Requests.ts";
import Label from "../../Components/Label.vue";
import Modal from "../../Components/Modal.vue";
import type { Energy } from "../../Models/Common/Energy.ts";

const mealStore = useMealStore();

const foodItemForm = ref<PostMealEntryRequest>({
    foodItemId: 0,
    amount: 0
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

function showFoodItemDialog(item: FoodItem) {
    selectFoodItem.value = false;
    foodItemForm.value.amount = 100;
    foodItemForm.value.foodItemId = item.id;
    selectedFoodItem.value = foodItemStore.getFoodItem(item.id);
}

const nutrients = computed((): Energy => {
    if (selectedFoodItem.value) {
        return {
            kCal: (selectedFoodItem.value.kCal * foodItemForm.value.amount) / 100,
            protein: (selectedFoodItem.value.protein * foodItemForm.value.amount) / 100,
            carbohydrate: (selectedFoodItem.value.carbohydrate * foodItemForm.value.amount) / 100,
            fat: (selectedFoodItem.value.fat * foodItemForm.value.amount) / 100
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
    <template v-if="selectFoodItem">
        <FormField>
            <InputText v-model="searchTerm" placeholder="Search"></InputText>
        </FormField>
        <ul>
            <li
                v-for="foodItem in foodItemStore.filteredFoodItems"
                :key="foodItem.id"
                @click="showFoodItemDialog(foodItem)"
                class="is-flex box is-justify-content-space-between"
            >
                <p><b>{{ foodItem.name }}</b></p>
                <p class="pr-2"><b>KCal:</b> {{ foodItem.kCal }}</p>
            </li>
        </ul>
    </template>
    <template v-else>
        <Modal @closed="() => selectFoodItem = true" :title="selectedFoodItem?.name">
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
                <Button @click="addToMeal">Add</Button>
            </template>
        </Modal>
    </template>
</template>

<style scoped></style>

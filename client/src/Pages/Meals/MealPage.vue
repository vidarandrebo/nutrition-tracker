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
import { useUserStore } from "../../Stores/UserStore.ts";
import { HttpRequest } from "http-methods-ts";
import type { MealEntryResponse, MealResponse } from "../../Models/Meals/Responses.ts";
import { MealEntry } from "../../Models/Meals/MealEntry.ts";

const mealStore = useMealStore();
const userStore = useUserStore();

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

function showFoodItemDialog(item: FoodItem) {
    selectFoodItem.value = false;
    foodItemForm.value.amount = 100;
    foodItemForm.value.foodItemId = item.id
    selectedFoodItem.value = foodItemStore.getFoodItem(item.id);
}

const nutrients = computed(() => {
    if (selectedFoodItem.value) {
        return {
            protein: (selectedFoodItem.value.protein * foodItemForm.value.amount) / 100,
            carb: (selectedFoodItem.value.carbohydrate * foodItemForm.value.amount) / 100,
            fat: (selectedFoodItem.value.fat * foodItemForm.value.amount) / 100,
        };
    }
    return { protein: 0, carb: 0, fat: 0 };
});

async function addToMeal() {
    if (meal.value) {
        await mealStore.addMealEntry(foodItemForm.value, meal.value.id)
    }
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
                class="is-flex"
            >
                <p>
                    {{ foodItem.name }}
                </p>
                <Button>+</Button>
            </li>
        </ul>
    </template>
    <template v-else>
        <div class="is-flex">
            <p>Name: {{ selectedFoodItem?.name }}</p>
            <Label>
                <p>Amount (g)</p>
                <p v-if="selectedFoodItem">Protein: {{ nutrients.protein }}</p>
                <p v-if="selectedFoodItem">Carb: {{ nutrients.carb }}</p>
                <p v-if="selectedFoodItem">Fat: {{ nutrients.fat }}</p>
                <InputNumber v-model="foodItemForm.amount"></InputNumber>
            </Label>
            <Button @click="addToMeal">Add</Button>
        </div>
    </template>
</template>

<style scoped></style>

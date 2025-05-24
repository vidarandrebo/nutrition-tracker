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
import type { Macronutrients } from "../../Models/Common/Macronutrients.ts";

const mealStore = useMealStore();

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

const nutrients = computed((): Macronutrients => {
    if (selectedFoodItem.value) {
        return {
            protein: (selectedFoodItem.value.protein * foodItemForm.value.amount) / 100,
            carbohydrate: (selectedFoodItem.value.carbohydrate * foodItemForm.value.amount) / 100,
            fat: (selectedFoodItem.value.fat * foodItemForm.value.amount) / 100,
        };
    }
    return { protein: 0, carbohydrate: 0, fat: 0 };
});

async function addToMeal() {
    if (meal.value) {
        await mealStore.addMealEntry(foodItemForm.value, meal.value.id)
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
        <Modal @closed="() => selectFoodItem = true" title="Add Item">
            <template #default>
                <div class="is-flex is-flex-direction-column">
                    <p><b>Name:</b> {{ selectedFoodItem?.name }}</p>
                    <p><b>Protein:</b> {{ nutrients.protein }} g</p>
                    <p><b>Carb:</b> {{ nutrients.carbohydrate }} g</p>
                    <p><b>Fat:</b> {{ nutrients.fat }} g</p>
                    <Label>
                        <p>Amount (g)</p>
                        <InputNumber v-model="foodItemForm.amount"></InputNumber>
                    </Label>
                </div>
            </template>
            <template #footer>
                <Button @click="addToMeal">Add</Button>
            </template>
        </Modal>
    </template>
</template>

<style scoped></style>

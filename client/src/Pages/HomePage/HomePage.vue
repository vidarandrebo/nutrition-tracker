<script setup lang="ts">
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useUserStore } from "../../Stores/UserStore.ts";
import InputDate from "../../Components/Forms/InputDate.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import { useMealStore } from "../../Stores/MealStore.ts";
import { computed, onMounted, watch } from "vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import { useMealViewStore } from "../../Stores/MealViewStore.ts";
import { addDays } from "../../Utilities/Date.ts";
import { MealEntry } from "../../Models/Meals/MealEntry.ts";
import { Meal } from "../../Models/Meals/Meal.ts";
import MealView from "./MealView.vue";

const userStore = useUserStore();
const mealStore = useMealStore();
const foodItemStore = useFoodItemStore();
const mealViewStore = useMealViewStore();

const foodItemIds = computed(() => [
    ...new Set(mealStore.mealsForDay.flatMap((m) => m.entries).map((f) => f.foodItemId)),
]);
onMounted(async () => {
    await mealStore.loadMealsForDay();
    await mealViewStore.init();
});
watch(foodItemIds, async () => {
    for (const id of foodItemIds.value) {
        if (id && !foodItemStore.getFoodItem(id)) {
            await FoodItem.getById(id).then((f) => {
                if (f) {
                    foodItemStore.collection.push(f);
                }
            });
        }
    }
});

function bumpDay(n: number) {
    mealStore.selectedDay = addDays(mealStore.selectedDay, n);
}

async function onDeleteEntry(entryId: number, mealId: number) {
    const { error } = await MealEntry.delete(entryId, mealId);
    if (error) {
        console.log("failed to delete meal entry");
        return;
    }
    const { error: rmEntryErr } = mealStore.removeMealEntry(entryId, mealId);
    if (rmEntryErr) {
        console.log(rmEntryErr.message);
    }
}

async function onDeleteMeal(mealId: number) {
    const { error } = await Meal.delete(mealId);
    if (error) {
        return;
    }

    const { error: rmMealErr } = mealStore.removeMeal(mealId);
    if (rmMealErr) {
        console.log(rmMealErr.message);
    }
}
</script>
<template>
    <section class="container">
        <HeaderH1>Home</HeaderH1>
        <div v-if="userStore.user">
            <div>
                <p>
                    KCal: {{ mealViewStore.dailyMacros.KCal }}, Protein: {{ mealViewStore.dailyMacros.Protein }} g,
                    Carbohydrate: {{ mealViewStore.dailyMacros.Carbohydrate }} g, Fat:
                    {{ mealViewStore.dailyMacros.Fat }} g
                </p>
            </div>
            <div class="is-flex is-justify-content-space-between">
                <div class="">
                    <div class="is-flex">
                        <div class="">
                            <ButtonPrimary @click="bumpDay(-1)">&larr;</ButtonPrimary>
                        </div>
                        <div class="">
                            <InputDate v-model="mealStore.selectedDay"></InputDate>
                        </div>
                        <div class="">
                            <ButtonPrimary @click="bumpDay(1)">&rarr;</ButtonPrimary>
                        </div>
                    </div>
                </div>
                <div class="is-narrow">
                    <ButtonPrimary @click="mealStore.addMeal">Add meal</ButtonPrimary>
                </div>
            </div>
        </div>
        <ul class="">
            <li v-for="item in mealViewStore.mealsView" :key="item.id" class="box">
                <MealView :item="item" @delete-meal="onDeleteMeal" @delete-meal-entry="onDeleteEntry"></MealView>
            </li>
        </ul>
    </section>
</template>

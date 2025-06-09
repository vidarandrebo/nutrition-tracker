<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import InputDate from "../Components/Forms/InputDate.vue";
import ButtonPrimary from "../Components/Buttons/ButtonPrimary.vue";
import { useMealStore } from "../Stores/MealStore.ts";
import { computed, onMounted, watch } from "vue";
import { useFoodItemStore } from "../Stores/FoodItemStore.ts";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";
import { useMealViewStore } from "../Stores/MealViewStore.ts";
import { addDays } from "../Utilities/Date.ts";

const userStore = useUserStore();
const mealStore = useMealStore();
const foodItemStore = useFoodItemStore();
const mealViewStore = useMealViewStore();

const foodItemIds = computed(() => [
    ...new Set(mealStore.mealsForDay.flatMap((m) => m.entries).map((f) => f.foodItemId)),
]);
onMounted(async () => {
    await mealStore.loadMealsForDay();
});
watch(foodItemIds, () => {
    for (const id of foodItemIds.value) {
        if (id && !foodItemStore.getFoodItem(id)) {
            FoodItem.getById(id).then((f) => {
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
</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <div>
            <p>
                KCal: {{ mealViewStore.dailyMacros.kCal }}, Protein: {{ mealViewStore.dailyMacros.protein }} g,
                Carbohydrate: {{ mealViewStore.dailyMacros.carbohydrate }} g, Fat: {{ mealViewStore.dailyMacros.fat }} g
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
                <ButtonPrimary v-on:click="mealStore.addMeal">Add meal</ButtonPrimary>
            </div>
        </div>
    </div>
    <ul class="">
        <li v-for="item in mealViewStore.mealsView" :key="item.id" class="box">
            <div>
                <RouterLink :to="{ path: '/meals/' + item.id }">Meal {{ item.id }}</RouterLink>
            </div>
            <div>
                {{ item.timestamp }}
            </div>
            <ul class="content">
                <li v-for="entry in item.entries" :key="entry.id" class="box">
                    <p>{{ entry.name }}, {{ entry.amount }}g.</p>
                    <p>
                        KCal: {{ entry.kCal }}, Protein: {{ entry.protein }}, Carbohydrate: {{ entry.carbohydrate }},
                        Fat: {{ entry.fat }}
                    </p>
                </li>
            </ul>
        </li>
    </ul>
</template>

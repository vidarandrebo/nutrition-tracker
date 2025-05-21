<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import InputDate from "../Components/InputDate.vue";
import Button from "../Components/Button.vue";
import { useMealStore } from "../Stores/MealStore.ts";
import { computed, onMounted, watch } from "vue";
import { useFoodItemStore } from "../Stores/FoodItemStore.ts";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";
import type { MealView } from "../Models/Meals/MealView.ts";

const userStore = useUserStore();
const mealStore = useMealStore();
const foodItemStore = useFoodItemStore();


const foodItemIds = computed(() => [...new Set(mealStore.mealsForDay
    .map((m) => m.entries)
    .flat(1)
    .map((f) => f.foodItemId))
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

const mealsView = computed((): MealView[] => {
    return mealStore.mealsForDay.map((m) => {
        return {
            id: m.id,
            timestamp: m.timestamp,
            entries: m.entries.map((me) => {
                const fi = foodItemStore.getFoodItem(me.foodItemId);
                return {
                    id: me.id,
                    name: fi?.name ?? "",
                    protein: fi ? (fi.protein * me.amount) / 100 : 0.0,
                    carbohydrate: fi ? (fi.carbohydrate * me.amount) / 100 : 0.0,
                    fat: fi ? (fi.fat * me.amount) / 100 : 0.0,
                    kCal: fi ? (fi.kCal * me.amount) / 100 : 0.0
                };
            })
        };
    });
});

</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <p>Welcome back {{ userStore.user.email }}</p>
        <InputDate v-model="mealStore.selectedDay"></InputDate>
    </div>
    <Button v-on:click="mealStore.addMeal">Add meal</Button>
    <Button v-on:click="mealStore.loadMealsForDay">Get meals</Button>
    <ul>
        <li v-for="item in mealsView" :key="item.id" class="box">
            <div>
                <RouterLink :to="{path: '/meals/' + item.id}">{{ item.id }}</RouterLink>
            </div>
            <div>
                {{ item.timestamp }}
            </div>
            <div v-for="entry in item.entries">
                {{ entry.name }} {{ entry.protein }}
            </div>
        </li>
    </ul>
</template>

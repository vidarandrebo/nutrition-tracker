<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import InputDate from "../Components/InputDate.vue";
import Button from "../Components/Button.vue";
import { useMealStore } from "../Stores/MealStore.ts";
import { computed, onMounted, watch } from "vue";
import { useFoodItemStore } from "../Stores/FoodItemStore.ts";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";
import { useMealViewStore } from "../Stores/MealViewStore.ts";

const userStore = useUserStore();
const mealStore = useMealStore();
const foodItemStore = useFoodItemStore();
const mealViewStore = useMealViewStore();


const foodItemIds = computed(() => [...new Set(mealStore.mealsForDay
    .flatMap((m) => m.entries)
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


</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <p>Welcome back {{ userStore.user.email }}</p>
        <InputDate v-model="mealStore.selectedDay"></InputDate>
    </div>
    <Button v-on:click="mealStore.addMeal">Add meal</Button>
    <Button v-on:click="mealStore.loadMealsForDay">Get meals</Button>
    <ul class="">
        <li v-for="item in mealViewStore.mealsView" :key="item.id" class="box">
            <div>
                <RouterLink :to="{path: '/meals/' + item.id}">{{ item.id }}</RouterLink>
            </div>
            <div>
                {{ item.timestamp }}
            </div>
            <ul class="content">
                <li v-for="entry in item.entries">
                    <p>{{ entry.name }}, {{ entry.amount }}g.</p>
                    <p>KCal: {{ entry.kCal }}, Protein: {{ entry.protein }}, Carbohydrate: {{ entry.carbohydrate }},
                        Fat: {{ entry.fat }}</p>
                </li>
            </ul>
        </li>
    </ul>
</template>

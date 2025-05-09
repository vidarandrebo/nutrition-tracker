<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import { ref, watch } from "vue";
import InputDate from "../Components/InputDate.vue";
import Button from "../Components/Button.vue";
import { Meal } from "../Models/Meals/Meal.ts";
import { useMealStore } from "../Stores/MealStore.ts";

const userStore = useUserStore();

const mealStore = useMealStore();

const selectedDay = ref(new Date());

watch(selectedDay, (newValue, oldValue) => {
    console.log(`Day changed from ${oldValue} to ${newValue}`);
});

async function getMeals() {
   const meals =  await Meal.get(selectedDay.value);
   if (meals) {
       mealStore.collection.push(...meals)
   }
}
function addMeal() {
    Meal.add(selectedDay.value)
}

</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <p>Welcome back {{ userStore.user.email }}</p>
        <InputDate v-model="selectedDay"></InputDate>
    </div>
    <Button v-on:click="addMeal">Add meal</Button>
    <Button v-on:click="getMeals">Get meals</Button>
    <ul>
        <li v-for="item in mealStore.collection" class="box">
            <div>
                {{item.id}}
            </div>
            <div>
                {{item.timestamp}}
            </div>
        </li>
    </ul>
</template>

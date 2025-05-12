<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import InputDate from "../Components/InputDate.vue";
import Button from "../Components/Button.vue";
import { useMealStore } from "../Stores/MealStore.ts";
import { onMounted } from "vue";

const userStore = useUserStore();

onMounted(async () => {
    await mealStore.loadMeals()
})

const mealStore = useMealStore();
</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <p>Welcome back {{ userStore.user.email }}</p>
        <InputDate v-model="mealStore.selectedDay"></InputDate>
    </div>
    <Button v-on:click="mealStore.addMeal">Add meal</Button>
    <Button v-on:click="mealStore.loadMeals">Get meals</Button>
    <ul>
        <li v-for="item in mealStore.mealsForDay" :key="item.id" class="box">
            <div>
                {{ item.id }}
            </div>
            <div>
                {{ item.timestamp }}
            </div>
        </li>
    </ul>
</template>

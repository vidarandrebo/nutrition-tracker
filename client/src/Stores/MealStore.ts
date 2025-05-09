import { defineStore } from "pinia";
import { ref } from "vue";
import { Meal } from "../Models/Meals/Meal.ts";

export const useMealStore = defineStore("meals", () => {
    const collection = ref<Meal[]>([]);

    return { collection };
});
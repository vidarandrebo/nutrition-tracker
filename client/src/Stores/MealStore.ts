import { defineStore } from "pinia";
import { computed, ref, watch } from "vue";
import { Meal } from "../Models/Meals/Meal.ts";
import { addDays, startOfDay } from "../Utilities/Date.ts";

export const useMealStore = defineStore("meals", () => {
    const selectedDay = ref<Date>(new Date());
    const collection = ref<Meal[]>([]);

    watch(selectedDay, loadMeals);

    const mealsForDay = computed(() => {
        const startTs = startOfDay(selectedDay.value);
        const endTs = addDays(startTs, 1)

        return collection.value.filter((m) => m.timestamp >= startTs && m.timestamp < endTs)
    });
    async function loadMeals() {
        const meals =  await Meal.get(selectedDay.value);
        meals?.map((m) => {
            if (collection.value.find((x) => x.id === m.id) === undefined) {
                collection.value.push(m)
            }
        })
    }

    async function addMeal() {
        const meal = await Meal.add(selectedDay.value)
        if (meal) {
            collection.value.push(meal)
        }
    }

    return { collection, loadMeals, mealsForDay, selectedDay, addMeal};
});
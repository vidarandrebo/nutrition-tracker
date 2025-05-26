import { defineStore } from "pinia";
import { computed, ref, watch } from "vue";
import { Meal } from "../Models/Meals/Meal.ts";
import { addDays, startOfDay } from "../Utilities/Date.ts";
import type { PostMealEntryRequest } from "../Models/Meals/Requests.ts";
import { MealEntry } from "../Models/Meals/MealEntry.ts";

export const useMealStore = defineStore("meals", () => {
    const selectedDay = ref<Date>(new Date());
    const collection = ref<Meal[]>([]);

    watch(selectedDay, loadMealsForDay);

    function clear() {
        collection.value = [];
    }

    const mealsForDay = computed(() => {
        const startTs = startOfDay(selectedDay.value);
        const endTs = addDays(startTs, 1)

        return collection.value.filter((m) => m.timestamp >= startTs && m.timestamp < endTs)
    });
    async function loadMealsForDay() {
        const meals =  await Meal.getByDay(selectedDay.value);
        meals?.map((m) => {
            if (collection.value.find((x) => x.id === m.id) === undefined) {
                collection.value.push(m)
            }
        })
    }

    function getMeal(id: number):  Meal | null {
        const meal = collection.value.find((m) => m.id === id)
        if (meal) {
            return meal
        }

        return null
    }

    async function loadMeal(id: number) {
        const meal = await Meal.getById(id);
        if (meal) {
            collection.value.push(meal)
        }
    }

    async function addMeal() {
        const meal = await Meal.add(selectedDay.value)
        if (meal) {
            collection.value.push(meal)
        }
    }

    async function addMealEntry(entry: PostMealEntryRequest, mealID: number) {
        const newEntry = await MealEntry.add(entry, mealID)
        if (newEntry) {
            const meal = collection.value.find((m) => m.id === mealID);
            if (meal) {
                meal.entries.push(newEntry)
            }
        }
    }

    return { clear, collection, loadMealsForDay, mealsForDay, selectedDay, addMeal, addMealEntry, getMeal, loadMeal};
});
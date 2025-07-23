import { defineStore } from "pinia";
import { computed, ref, watch } from "vue";
import { Meal } from "../Models/Meals/Meal.ts";
import { addDays, startOfDay } from "../Utilities/Date.ts";
import type { PostMealEntryRequest } from "../Models/Meals/Requests.ts";
import { MealEntry } from "../Models/Meals/MealEntry.ts";
import { Failure, type Result, Success } from "../Utilities/tryCatch.ts";

export const useMealStore = defineStore("meals", () => {
    const selectedDay = ref<Date>(new Date());
    const collection = ref<Meal[]>([]);

    watch(selectedDay, loadMealsForDay);

    function clear() {
        collection.value = [];
    }

    const mealsForDay = computed(() => {
        const startTs = startOfDay(selectedDay.value);
        const endTs = addDays(startTs, 1);

        return collection.value.filter((m) => m.timestamp >= startTs && m.timestamp < endTs);
    });

    async function loadMealsForDay() {
        const meals = await Meal.getByDay(selectedDay.value);
        meals?.map((m) => {
            if (collection.value.find((x) => x.id === m.id) === undefined) {
                collection.value.push(m);
            }
        });
    }

    function getMeal(id: number): Meal | null {
        const meal = collection.value.find((m) => m.id === id);
        if (meal) {
            return meal;
        }

        return null;
    }

    async function loadMeal(id: number) {
        const meal = await Meal.getById(id);
        if (meal) {
            collection.value.push(meal);
        }
    }

    async function addMeal() {
        const meal = await Meal.add(selectedDay.value);
        if (meal) {
            collection.value.push(meal);
        }
    }

    async function addMealEntry(entry: PostMealEntryRequest, mealID: number) {
        const newEntry = await MealEntry.add(entry, mealID);
        if (newEntry) {
            const meal = collection.value.find((m) => m.id === mealID);
            if (meal) {
                meal.entries.push(newEntry);
            }
        }
    }

    function removeMealEntry(entryId: number, mealId: number): Result<void> {
        const entries = collection.value.find((x) => (x.id = mealId))?.entries;
        if (entries) {
            const idx = entries.findIndex((x) => x.id === entryId);
            if (idx !== -1) {
                entries?.splice(idx, 1);
                return Success.empty();
            }
        }
        return Failure.new(new Error("failed to remove meal entry"));
    }

    function removeMeal(mealId: number): Result<void> {
        const idx = collection.value.findIndex((m) => m.id === mealId);
        if (idx !== -1) {
            collection.value.splice(idx, 1);
            return Success.empty();
        }
        return Failure.new(new Error("failed to remove meal"));
    }

    return {
        addMeal,
        addMealEntry,
        clear,
        collection,
        getMeal,
        loadMeal,
        loadMealsForDay,
        mealsForDay,
        removeMeal,
        removeMealEntry,
        selectedDay,
    };
});

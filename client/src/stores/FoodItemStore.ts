import { defineStore } from "pinia";
import { ref } from "vue";
import type { FoodItem } from "../models/Fooditem.ts";

export const useFoodItemStore = defineStore("foodItems", () => {
    const collection = ref<FoodItem[]>([])
    return {foodItems: collection}
})
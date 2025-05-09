import { defineStore } from "pinia";
import { ref } from "vue";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";

export const useFoodItemStore = defineStore("foodItems", () => {
    const collection = ref<FoodItem[]>([]);

    async function init() {
        if (collection.value.length === 0) {
            const items = await FoodItem.get();
            if (items === null) {
                collection.value = [];
            } else {
                collection.value = items;
            }
        }
    }

    async function refresh() {
        const items = await FoodItem.get();
        if (items === null) {
            collection.value = [];
        } else {
            collection.value = items;
        }
    }

    return { collection, init, refresh };
});

import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";

export const useFoodItemStore = defineStore("foodItems", () => {
    const collection = ref<FoodItem[]>([]);
    const initialized = ref<boolean>(false);
    const searchTerm = ref<string>("");

    function clear() {
        collection.value = [];
        initialized.value = false;
    }

    async function init() {
        searchTerm.value = "";
        if (!initialized.value) {
            const items = await FoodItem.get();
            if (items === null) {
                collection.value = [];
            } else {
                collection.value = items;
            }
            initialized.value = true;
        }
    }

    function getFoodItem(id: number): FoodItem | undefined {
        return collection.value.find((f) => f.id === id);
    }

    async function refresh() {
        const items = await FoodItem.get();
        if (items === null) {
            collection.value = [];
        } else {
            collection.value = items;
        }
    }

    const filteredFoodItems = computed(() => {
        if (searchTerm.value.length < 3) {
            return collection.value;
        }
        const terms = searchTerm.value
            .split(" ")
            .filter((s) => s !== "")
            .map((t) => t.toLowerCase());

        if (searchTerm.value === "") {
            return collection.value;
        }
        return collection.value
            .filter((x) => {
                for (let i = 0; i < terms.length; i++) {
                    if (!x.name.toLowerCase().includes(terms[i])) {
                        return false;
                    }
                }
                return true;
            })
            .sort((a, b) => {
                if (a.product.length < b.product.length) {
                    return -1;
                } else if (a.product.length > b.product.length) {
                    return 1;
                } else {
                    return 0;
                }
            })
            .slice(0, 50);
    });
    return { clear, collection, init, refresh, filteredFoodItems, searchTerm, getFoodItem, initialized };
});

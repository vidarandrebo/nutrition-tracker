import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { FoodItem } from "../Models/FoodItems/Fooditem.ts";

export const useFoodItemStore = defineStore("foodItems", () => {
    const collection = ref<FoodItem[]>([]);
    const initialized = ref<boolean>(false)

    function clear() {
        collection.value = [];
        initialized.value = false;
    }

    async function init() {
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

    function getFoodItem(id: number) : FoodItem | undefined {
        return collection.value.find((f) => f.id === id)
    }

    async function refresh() {
        const items = await FoodItem.get();
        if (items === null) {
            collection.value = [];
        } else {
            collection.value = items;
        }
    }

    const searchTerm = ref<string>("");
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
        return collection.value.filter((x) => {
            for (let i = 0; i < terms.length; i++) {
                if (!(x.product.toLowerCase().includes(terms[i]) || x.manufacturer.toLowerCase().includes(terms[i]))) {
                    return false;
                }
            }
            return true;
        })
            .slice(0,25);
    });
    return { clear, collection, init, refresh , filteredFoodItems, searchTerm, getFoodItem, initialized};
});

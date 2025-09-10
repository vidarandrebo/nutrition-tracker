import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { Failure, type Result, Success } from "../Utilities/tryCatch.ts";
import { FoodItem } from "../Models/FoodItems/FoodItem.ts";
import { PortionSize, type PortionSizeForm } from "../Models/FoodItems/PortionSize.ts";
import { useFilterStore } from "./FilterStore.ts";
import { useUserStore } from "./UserStore.ts";

export const useFoodItemStore = defineStore("foodItems", () => {
    const collection = ref<FoodItem[]>([]);
    const initialized = ref<boolean>(false);
    const searchTerm = ref<string>("");
    const filterStore = useFilterStore();
    const userStore = useUserStore();
    filterStore.init();

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
    function removeFoodItem(id: number): Result<void> {
        const idx = collection.value.findIndex((m) => m.id === id);
        if (idx !== -1) {
            collection.value.splice(idx, 1);
            return Success.empty();
        }
        return Failure.new(new Error("failed to remove foodItem"));
    }

    const filteredFoodItems = computed(() => {
        const items = filterStore.foodItem.applyFilter(collection.value, {
            ownerId: userStore.user?.id ?? 0,
            searchTerm: searchTerm.value,
        });
        return items
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
    async function addPortionSize(foodItemId: number, ps: PortionSizeForm) {
        const result = await PortionSize.add(ps, foodItemId);
        if (!result.error) {
            const data = result.data;
            const portionSize = PortionSize.fromResponse(data);
            const fi = collection.value.find((f) => f.id === foodItemId);
            if (fi) {
                fi.portionSizes.push(portionSize);
            }
        }
    }
    return {
        addPortionSize,
        clear,
        collection,
        filteredFoodItems,
        getFoodItem,
        init,
        initialized,
        refresh,
        removeFoodItem,
        searchTerm,
    };
});

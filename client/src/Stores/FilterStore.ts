import { defineStore } from "pinia";
import { ref } from "vue";
import { type FoodItem } from "../Models/FoodItems/FoodItem.ts";

interface IFilter<T> {
    applyFilter(elements: T[]): T[];
}

class FoodItemFilter implements IFilter<FoodItem> {
    showPublic: boolean;
    constructor() {
        this.showPublic = false;
    }

    applyFilter(elements: FoodItem[]): FoodItem[] {
        return elements.filter((fi) => {
            if (this.)
            if (this.showPublic) {
                return true;
            } else if (fi.public {
            }
        });
    }
    static fromJson(s: string): FoodItemFilter {
        const filter = new FoodItemFilter();
        const o = JSON.parse(s);

        const showPublic = o["showPublic"];
        if (typeof showPublic === "boolean") {
            filter.showPublic = showPublic;
        }

        return filter;
    }
}

export const useFilterStore = defineStore("filterStore", () => {
    const foodItem = ref<FoodItemFilter>(new FoodItemFilter());

    function init() {
        loadFoodItemFilter();
    }

    function setFoodItemFilter(foodItemFilter: FoodItemFilter) {
        foodItem.value = foodItemFilter;
        localStorage.setItem("foodItemFilter", JSON.stringify(foodItem.value));
    }
    function loadFoodItemFilter() {
        const f = localStorage.getItem("foodItemFilter");
        if (f) {
            foodItem.value = FoodItemFilter.fromJson(f);
        }
    }
    return {
        init,
        setFoodItemFilter,
        loadFoodItemFilter,
        foodItem,
    };
});

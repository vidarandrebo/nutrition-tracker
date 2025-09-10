import { defineStore } from "pinia";
import { ref } from "vue";
import { type FoodItem } from "../Models/FoodItems/FoodItem.ts";

interface IFilter<T> {
    applyFilter(elements: T[], filterOptions: never): T[];
}

type FoodItemFilterOptions = {
    ownerId: number;
    searchTerm: string;
};

class FoodItemFilter implements IFilter<FoodItem> {
    showPublic: boolean;
    constructor() {
        this.showPublic = false;
    }

    applyFilter(elements: FoodItem[], filterOptions: FoodItemFilterOptions): FoodItem[] {
        const terms = filterOptions.searchTerm
            .split(" ")
            .filter((s) => s !== "")
            .map((t) => t.toLowerCase());

        console.log(terms);
        console.log(filterOptions);
        return elements
            .filter((fi) => {
                if (filterOptions.searchTerm.length < 3) {
                    return true;
                }
                for (let i = 0; i < terms.length; i++) {
                    const term = terms[i];
                    if (!term) {
                        continue;
                    }
                    if (!fi.name.toLowerCase().includes(term)) {
                        return false;
                    }
                }
                return true;
            })
            .filter((fi) => {
                if (fi.ownerId === filterOptions.ownerId) {
                    return true;
                }
                if (this.showPublic) {
                    return true;
                }
                return false;
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
    const loaded = ref<boolean>(false);
    const foodItem = ref<FoodItemFilter>(new FoodItemFilter());

    function init() {
        if (loaded.value) {
            loadFoodItemFilter();
        }
        loaded.value = true;
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

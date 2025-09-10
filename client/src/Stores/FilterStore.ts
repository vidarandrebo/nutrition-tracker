import { defineStore } from "pinia";
import { ref, watch } from "vue";
import { type FoodItem } from "../Models/FoodItems/FoodItem.ts";

interface IFilter<T> {
    applyFilter(elements: T[], filterOptions: never): T[];
}

type FoodItemFilterOptions = {
    ownerId: number;
};

class FoodItemFilter implements IFilter<FoodItem> {
    showPublic: boolean;
    searchTerm: string;
    constructor() {
        this.searchTerm = "";
        this.showPublic = false;
    }

    applyFilter(elements: FoodItem[], filterOptions: FoodItemFilterOptions): FoodItem[] {
        const terms = this.searchTerm
            .split(" ")
            .filter((s) => s !== "")
            .map((t) => t.toLowerCase());

        return elements
            .filter((fi) => {
                if (this.searchTerm.length < 3) {
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
    writeToLocalStorage() {
        localStorage.setItem("foodItemFilter", JSON.stringify(this));
    }
    static readFromLocalStorage(): FoodItemFilter | null {
        const f = localStorage.getItem("foodItemFilter");
        if (!f) {
            return null;
        }
        return this.fromJson(f);
    }
}

export const useFilterStore = defineStore("filter", () => {
    const foodItem = ref<FoodItemFilter>(FoodItemFilter.readFromLocalStorage() ?? new FoodItemFilter());

    watch(
        foodItem,
        () => {
            foodItem.value.writeToLocalStorage();
        },
        { deep: true },
    );
    return {
        foodItem,
    };
});

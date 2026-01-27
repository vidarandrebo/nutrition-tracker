<script setup lang="ts">
import { EntryType, type MealView } from "../../Models/Meals/MealView.ts";
import { ref } from "vue";
import { OnClickOutside } from "@vueuse/components";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import MealEntryView from "./MealEntryView.vue";

const props = defineProps<{
    item: MealView;
}>();
const emit = defineEmits<{
    deleteMeal: [id: number];
    deleteMealFoodItemEntry: [entryId: number, mealId: number];
    deleteMealRecipeEntry: [entryId: number, mealId: number];
    deleteMealMacronutrientEntry: [entryId: number, mealId: number];
}>();

const kebabOpen = ref<boolean>(false);

function onClickOutsideHandler() {
    kebabOpen.value = false;
}
function onDeleteMealEntry(entryId: number, mealId: number, entryType: EntryType) {
    switch (entryType) {
        case EntryType.FoodItem:
            console.debug("deleting food item entry from meal", entryId, mealId);
            emit("deleteMealFoodItemEntry", entryId, mealId);
            break;
        case EntryType.Macronutrient:
            console.debug("deleting macronutrient entry from meal", entryId, mealId);
            emit("deleteMealMacronutrientEntry", entryId, mealId);
            break;
        case EntryType.Recipe:
            console.debug("deleting recipe entry from meal", entryId, mealId);
            emit("deleteMealRecipeEntry", entryId, mealId);
            break;
        default:
            break;
    }
}
</script>

<template>
    <LevelPrimary>
        <template #left>
            <div>
                <RouterLink :to="{ path: '/meals/' + props.item.id }">Meal {{ props.item.id }}</RouterLink>
            </div>
        </template>
        <template #right>
            <OnClickOutside @trigger="onClickOutsideHandler">
                <div :class="kebabOpen ? 'dropdown is-active is-right' : 'dropdown is-right'">
                    <div class="dropdown-trigger">
                        <button
                            class="button"
                            aria-haspopup="true"
                            aria-controls="dropdown-menu"
                            @click="() => (kebabOpen = !kebabOpen)"
                        >
                            <span>⋮</span>
                        </button>
                    </div>
                    <div class="dropdown-menu" role="menu">
                        <div class="dropdown-content">
                            <a href="#" class="dropdown-item" @click="() => emit('deleteMeal', props.item.id)"
                                >Delete meal</a
                            >
                        </div>
                    </div>
                </div>
            </OnClickOutside>
        </template>
    </LevelPrimary>
    <div>
        {{ props.item.timestamp }}
    </div>
    <ul class="content">
        <li v-for="entry in props.item.entries" :key="entry.id" class="box">
            <MealEntryView
                :entry="entry"
                @delete-meal-entry="() => onDeleteMealEntry(entry.id, props.item.id, entry.entryType)"
            ></MealEntryView>
        </li>
    </ul>
</template>
<style scoped></style>

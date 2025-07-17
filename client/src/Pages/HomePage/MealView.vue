<script setup lang="ts">
import ButtonDanger from "../../Components/Buttons/ButtonDanger.vue";
import type { MealView } from "../../Models/Meals/MealView.ts";
import { ref } from "vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";

const props = defineProps<{
    item: MealView;
}>();
const emit = defineEmits<{
    deleteMeal: [id: number];
    deleteMealEntry: [entryId: number, mealId: number];
}>();

const kebabOpen = ref<boolean>(false);
</script>

<template>
    <LevelPrimary>
        <template #left>
            <div>
                <RouterLink :to="{ path: '/meals/' + props.item.id }">Meal {{ props.item.id }}</RouterLink>
            </div>
        </template>
        <template #right>
            <div :class="kebabOpen ? 'dropdown is-active is-hoverable is-right' : 'dropdown is-hoverable is-right'">
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
                <div id="dropdown-menu" class="dropdown-menu" role="menu">
                    <div class="dropdown-content">
                        <a href="#" class="dropdown-item" @click="() => emit('deleteMeal', props.item.id)">Delete</a>
                    </div>
                </div>
            </div>
        </template>
    </LevelPrimary>
    <div>
        {{ props.item.timestamp }}
    </div>
    <ul class="content">
        <li v-for="entry in props.item.entries" :key="entry.id" class="box">
            <p>{{ entry.name }}, {{ entry.amount }}</p>
            <p>
                KCal: {{ entry.KCal }}, Protein: {{ entry.Protein }}, Carbohydrate: {{ entry.Carbohydrate }}, Fat:
                {{ entry.Fat }}
            </p>
            <ButtonDanger @click="() => emit('deleteMealEntry', entry.id, props.item.id)">Delete</ButtonDanger>
        </li>
    </ul>
</template>
<style scoped></style>

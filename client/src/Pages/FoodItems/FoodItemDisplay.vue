<script setup lang="ts">
import type { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { OnClickOutside } from "@vueuse/components";
import { ref } from "vue";

type FoodItemDisplayProps = {
    item: FoodItem;
};
const emit = defineEmits<{
    deleteFoodItem: [id: number];
}>();
const props = defineProps<FoodItemDisplayProps>();
const kebabOpen = ref<boolean>(false);

function onClickOutsideHandler() {
    kebabOpen.value = false;
}
</script>

<template>
    <li class="box">
        <LevelPrimary>
            <template #left>
                <RouterLink class="subtitle is-4" :to="{ path: '/food-items/' + props.item.id }">
                    <b> {{ item.name }}</b>
                </RouterLink>
            </template>
            <template #right>
                <!-- TODO check ownership to ensure fooditem can be deleted-->
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
                                <a href="#" class="dropdown-item" @click="() => emit('deleteFoodItem', props.item.id)"
                                    >Delete food item</a
                                >
                            </div>
                        </div>
                    </div>
                </OnClickOutside>
            </template>
        </LevelPrimary>
        <p>Protein: {{ props.item.protein }}</p>
        <p>Carbohydrate: {{ props.item.carbohydrate }}</p>
        <p>Fat: {{ props.item.fat }}</p>
        <p>KCal: {{ props.item.kCal }}</p>
    </li>
</template>

<style scoped></style>

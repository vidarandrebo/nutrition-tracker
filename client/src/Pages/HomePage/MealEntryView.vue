<script setup lang="ts">
import type { MealEntryView } from "../../Models/Meals/MealView.ts";
import { OnClickOutside } from "@vueuse/components";
import { ref } from "vue";

const props = defineProps<{
    entry: MealEntryView;
}>();
const emit = defineEmits<{
    deleteMealEntry: [entryId: number];
}>();
const kebabOpen = ref<boolean>(false);

function onClickOutsideHandler() {
    kebabOpen.value = false;
}
</script>

<template>
    <div class="is-flex is-flex-direction-row is-justify-content-space-between">
        <div class="is-flex-grow-1">
            <p>
                <b>{{ props.entry.name }}, {{ props.entry.amount }}</b>
            </p>
            <div class="columns is-gapless">
                <div class="column">
                    <p>KCal:&nbsp;{{ props.entry.KCal }}</p>
                </div>
                <div class="column">
                    <p>Protein:&nbsp;{{ props.entry.Protein }}&nbsp;g</p>
                </div>
                <div class="column">
                    <p>Carbohydrate:&nbsp;{{ props.entry.Carbohydrate }}&nbsp;g</p>
                </div>
                <div class="column">
                    <p>Fat:&nbsp;{{ props.entry.Fat }}&nbsp;g</p>
                </div>
            </div>
        </div>
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
                        <a href="#" class="dropdown-item" @click="() => emit('deleteMealEntry', props.entry.id)"
                            >Delete entry</a
                        >
                    </div>
                </div>
            </div>
        </OnClickOutside>
    </div>
</template>

<style scoped></style>

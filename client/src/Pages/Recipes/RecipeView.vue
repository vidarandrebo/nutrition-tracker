<script setup lang="ts">
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { OnClickOutside } from "@vueuse/components";
import type { RecipeView } from "../../Models/Recipes/RecipeView.ts";
import { ref } from "vue";
const props = defineProps<{
    item: RecipeView;
}>();
const emit = defineEmits<{
    deleteRecipe: [id: number];
}>();
const kebabOpen = ref<boolean>(false);

function onClickOutsideHandler() {
    kebabOpen.value = false;
}
</script>

<template>
    <LevelPrimary>
        <template #left>
            <b>{{ item.name }}</b>
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
                            <a href="#" class="dropdown-item" @click="() => emit('deleteRecipe', props.item.id)"
                                >Delete recipe</a
                            >
                        </div>
                    </div>
                </div>
            </OnClickOutside>
        </template>
    </LevelPrimary>
    <p>
        KCal: {{ item.KCal }}, Protein: {{ item.Protein }} g, Carbohydrate: {{ item.Carbohydrate }} g, Fat:
        {{ item.Fat }} g
    </p>
</template>
<style scoped></style>

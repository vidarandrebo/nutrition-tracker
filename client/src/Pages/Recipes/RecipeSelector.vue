<script setup lang="ts">
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import { onMounted } from "vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import ButtonPlain from "../../Components/Buttons/ButtonPlain.vue";
import HeaderH2 from "../../Components/Headings/HeaderH2.vue";
import { useRecipeViewStore } from "../../Stores/RecipeViewStore";

const recipeViewStore = useRecipeViewStore();

const emit = defineEmits<{
    select: [id: number];
    cancel: [];
}>();

function cancel() {
    emit("cancel");
}

onMounted(async () => {
    await recipeViewStore.init();
});
</script>

<template>
    <section>
        <LevelPrimary>
            <template #left>
                <HeaderH2>Add Recipe</HeaderH2>
            </template>
            <template #right>
                <ButtonPlain class="level-item" @click="cancel">Cancel</ButtonPlain>
            </template>
        </LevelPrimary>
        <!--<LabelPrimary>
            Search (showing {{ recipeViewStore.recipesView.length }} entries)
            <InputText v-model="recipeViewStore.searchTerm" placeholder="Search"></InputText>
        </LabelPrimary>-->
        <div
            v-for="item in recipeViewStore.recipesView"
            :key="item.id"
            class="is-flex is-flex-direction-row is-justify-content-space-between box"
        >
            <p>
                <b>{{ item.name }}</b>
            </p>
            <p class="pr-2"><b>KCal:</b> {{ item.kCal }}</p>
            <ButtonPrimary @click="() => emit('select', item.id)">Add</ButtonPrimary>
        </div>
    </section>
</template>

<style scoped></style>

<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useRecipeStore } from "../../Stores/RecipeStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import router from "../../Router.ts";
import Level from "../../Components/Level.vue";
import { onMounted } from "vue";

const recipeStore = useRecipeStore();
onMounted(async () => {
    await recipeStore.init();
});

function addRecipe() {
    router.push("/recipes/add");
}
</script>

<template>
    <Level>
        <template #left>
            <HeaderH1 class="level-item">Recipes</HeaderH1>
        </template>
        <template #right>
            <ButtonPrimary @click="addRecipe" class="level-item">Add</ButtonPrimary>
        </template>
    </Level>
    <article v-for="item in recipeStore.collection" :key="item.id">
        {{ item.name }} {{ item.entries }}
    </article>
</template>

<style scoped></style>

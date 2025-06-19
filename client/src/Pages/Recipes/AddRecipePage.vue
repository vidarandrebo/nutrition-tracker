<script setup lang="ts">
import RecipeForm from "./RecipeForm.vue";
import HeaderH1 from "../../Components/HeaderH1.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import ButtonPlain from "../../Components/Buttons/ButtonPlain.vue";
import router from "../../Router.ts";

async function onCancel() {
    await router.push("/recipes");
}
</script>

<template>
    <LevelPrimary>
        <template #left>
            <HeaderH1>Add Recipe</HeaderH1>
        </template>
        <template #right>
            <ButtonPlain @click="onCancel">Cancel</ButtonPlain>
        </template>
    </LevelPrimary>
    <RecipeForm></RecipeForm>
</template>

<style scoped></style>

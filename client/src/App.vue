<script setup lang="ts">
import MenuBar from "./Components/MenuBar.vue";
</script>

<template>
    <MenuBar></MenuBar>
    <main>
        <RouterView></RouterView>
    </main>
    <footer></footer>
</template>

<style lang="scss">
@use "bulma/sass";
</style>

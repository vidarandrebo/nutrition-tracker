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
@use "bulma/sass"
    with(
        //$text: hsl(20, 100%, 10),
        //$link: hsl(20, 100%, 10),
        $primary: hsl(13, 90%, 58%),
        $info: hsl(208, 32%, 51%),
        $success: hsl(232, 27%, 35%),
        $warning: hsl(273, 36%, 21%),
        $danger: hsl(260, 35%, 10%),
        $size-1: 2rem,
        $size-2: 1.52rem,
        $size-3: 1.25rem,
        $size-4: 1rem,
        $size-5: 0.75rem,
        $size-6: 0.67rem,
        $size-7: 0.5rem,
        $family-primary:
            "Inter, SF Pro, Segoe UI, Roboto, Oxygen, Ubuntu, Helvetica Neue, Helvetica, Arial, sans-serif, IBM Plex",
        $family-secondary:
            "Inter, SF Pro, Segoe UI, Roboto, Oxygen, Ubuntu, Helvetica Neue, Helvetica, Arial, sans-serif, IBM Plex",
        $family-code: "Inconsolata, Hack, SF Mono, Roboto Mono, Source Code Pro, Ubuntu Mono, monospace, IBM Plex",
        $size-small: 0.75rem,
        $size-normal: 1rem,
        $size-medium: 1.25rem,
        $size-large: 1.5rem,
        $weight-light: 200,
        $weight-normal: 300,
        $weight-medium: 400,
        $weight-semibold: 500,
        $weight-bold: 600,
        $weight-extrabold: 700
    );
</style>

<template>
    <RouterLink class="menubar-router-link" :to="to">
        <slot></slot>
    </RouterLink>
</template>

<script setup lang="ts">
defineProps<{
    to: string;
}>();
</script>

<style scoped>
</style>
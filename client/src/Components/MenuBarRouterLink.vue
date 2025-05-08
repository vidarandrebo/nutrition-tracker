<template>
    <RouterLink class="navbar-item is-primary" :to="to">
        <slot></slot>
    </RouterLink>
</template>

<script setup lang="ts">
defineProps<{
    to: string;
}>();
</script>

<style scoped>
.router-link-active {
    background-color:lightblue;
}
</style>
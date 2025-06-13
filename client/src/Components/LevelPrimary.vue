<script setup lang="ts"></script>

<template>
    <div class="level is-mobile">
        <div class="level-left">
            <slot name="left"></slot>
        </div>
        <div class="level-right">
            <slot name="right"></slot>
        </div>
    </div>
</template>

<style scoped></style>

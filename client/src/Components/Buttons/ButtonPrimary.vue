<template>
    <button class="button is-primary" :disabled="!props.enabled" :title="!props.enabled ? props.disabledText : ''">
        <slot></slot>
    </button>
</template>
<script lang="ts" setup>
const props = withDefaults(
    defineProps<{
        enabled?: boolean;
        disabledText?: string;
    }>(),
    {
        enabled: true,
        disabledText: "",
    },
);
</script>

<style scoped></style>

<script setup lang="ts">
defineProps<{
    modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

function handleInput(event: Event) {
    const inputValue = (event.target as HTMLInputElement).value;
    emit("update:modelValue", inputValue);
}
</script>

<template>
    <input :value="modelValue" class="input" @input="handleInput" />
</template>

<style scoped></style>

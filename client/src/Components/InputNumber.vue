<script setup lang="ts">
defineProps<{
    modelValue: number | undefined;
    placeHolder?: number | undefined
}>();
const emit = defineEmits(["update:modelValue"]);

function handleInput(event: Event) {
    const inputValue = (event.target as HTMLInputElement).value;
    emit("update:modelValue", inputValue);
}
</script>

<template>
    <input type="number" class="input" :value="modelValue" @input="handleInput" :placeholder="placeHolder?.toString()"/>
</template>

<style scoped></style>

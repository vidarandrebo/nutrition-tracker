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
    <input type="number" :value="modelValue" @input="handleInput" class="w-15" :placeholder="placeHolder?.toString()"/>
</template>

<style scoped></style>

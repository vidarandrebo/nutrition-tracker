<script setup lang="ts">
defineProps<{
    modelValue: number | undefined;
    placeHolder?: number | undefined;
}>();
const emit = defineEmits(["update:modelValue"]);

function handleInput(event: Event) {
    const inputValue = (event.target as HTMLInputElement).value;
    emit("update:modelValue", parseFloat(inputValue));
}
</script>

<template>
    <input
        type="number"
        class="input"
        :value="modelValue"
        :placeholder="placeHolder?.toString()"
        @input="handleInput"
    />
</template>

<style scoped></style>

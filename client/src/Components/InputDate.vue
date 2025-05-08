<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    modelValue: Date;
}>();
const emit = defineEmits(["update:modelValue"]);

function getDateString(date: Date): string {
    const day = date.getDate().toString().padStart(2, "0");
    const month = (date.getMonth() + 1).toString().padStart(2, "0"); // 0-index here...
    const year = date.getFullYear().toString();
    return `${year}-${month}-${day}`;
}

const dateString = computed(() => {
    return getDateString(props.modelValue);
});

function handleInput(event: Event) {
    const inputValue = (event.target as HTMLInputElement).value;
    emit("update:modelValue", new Date(inputValue));
}
</script>

<template>
    <input type="date" :value="dateString" @input="handleInput"/>
</template>

<style scoped></style>

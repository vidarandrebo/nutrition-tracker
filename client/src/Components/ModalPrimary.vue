<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    modelValue: boolean;
    title?: string;
}>();
const emit = defineEmits<{
    (e: "update:modelValue", payload: boolean): void;
}>();

function close() {
    emit("update:modelValue", false);
}

const modalActiveClass = computed(() => {
    if (props.modelValue) {
        return "modal is-active";
    }
    return "modal";
});
</script>
<template>
    <div :class="modalActiveClass">
        <div class="modal-background" @click="close"></div>
        <div class="modal-card">
            <header class="modal-card-head">
                <h2 class="modal-card-title">{{ props.title }}</h2>
                <button class="delete" aria-label="close" @click="close"></button>
                <slot name="header"></slot>
            </header>
            <section class="modal-card-body">
                <!-- Content ... -->
                <slot name="default"></slot>
            </section>
            <footer class="modal-card-foot">
                <div class="buttons">
                    <slot name="footer"> </slot>
                    <button class="button" @click="close">Cancel</button>
                </div>
            </footer>
        </div>
    </div>
</template>

<style scoped></style>

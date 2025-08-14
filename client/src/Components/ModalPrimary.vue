<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
    title?: string;
}>();
const model = defineModel<boolean>();

function close() {
    model.value = false;
}

const modalActiveClass = computed(() => {
    if (model.value) {
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

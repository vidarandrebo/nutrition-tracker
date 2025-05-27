<script setup lang="ts">

import { computed, ref, watch } from "vue";

const props = defineProps<{
    title?: string;
}>();

const isActive = ref<boolean>(true);
const emits = defineEmits<{
    (e: "closed"): void
}>();

watch(isActive, (value) => {
    if (!value) {
        emits("closed");
    }
});

const modalActiveClass = computed(() => {
    if (isActive.value) {
        return "modal is-active";
    }
    return "modal";
});
</script>
<template>
    <div :class="modalActiveClass">
        <div class="modal-background" @click="() => isActive = false"></div>
        <div class="modal-card">
            <header class="modal-card-head">
                <h2 class="modal-card-title">{{ props.title }}</h2>
                <button class="delete" aria-label="close" @click="() => isActive = false"></button>
                <slot name="header"></slot>
            </header>
            <section class="modal-card-body">
                <!-- Content ... -->
                <slot name="default"></slot>
            </section>
            <footer class="modal-card-foot">
                <div class="buttons">
                    <slot name="footer">
                    </slot>
                    <button class="button" @click="() => isActive= false">Cancel</button>
                </div>
            </footer>
        </div>
    </div>
</template>

<style scoped>

</style>
<script setup lang="ts">
import { ref } from "vue";

const props = defineProps<{
    entries: string[];
    preselected: string;
}>();

const emit = defineEmits<{
    selected: [tab: string];
}>();
const selected = ref<string>(props.preselected);
function selectTab(t: string) {
    selected.value = t;
    emit("selected", t);
}
</script>

<template>
    <div class="tabs">
        <ul>
            <template v-for="entry in props.entries" :key="entry">
                <li :class="selected === entry ? 'is-active' : null">
                    <a @click="() => selectTab(entry)">{{ entry }}</a>
                </li>
            </template>
        </ul>
    </div>
</template>

<style scoped></style>

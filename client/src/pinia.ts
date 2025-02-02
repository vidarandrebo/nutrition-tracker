import { createPinia, defineStore } from "pinia";
import { computed, ref } from "vue";

export const useCountStore = defineStore("countStore", () => {
    const count = ref<number>(0);

    const getCount = computed(() => {
        return count;
    });

    function setCount(v: number) {
        count.value = v;
    }
    return { getCount, setCount: setCount };
});

export const pinia = createPinia();

import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("user", () => {
    const token = ref<string>("")
    const id = ref<number>(0)

    function load() {
    }

    return {id, token, load}
})

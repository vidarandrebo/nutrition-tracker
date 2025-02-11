import { defineStore } from "pinia";
import { ref } from "vue";
import { readFromLocalStorage, type User } from "../Models/User.ts";

export const useUserStore = defineStore("user", () => {
    const user = ref<User | null>(readFromLocalStorage());

    return { user };
});

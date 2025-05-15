import { defineStore } from "pinia";
import { ref } from "vue";
import { User } from "../Models/User.ts";

export const useUserStore = defineStore("user", () => {
    const user = ref<User | null>(User.readFromLocalStorage());

    return { user };
});

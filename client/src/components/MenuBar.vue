<script setup lang="ts">
import { useUserStore } from "../stores/UseUserStore.ts";
import router from "../router.ts";

const userStore = useUserStore();

async function logout() {
    userStore.user = null

    localStorage.removeItem("user")

    await router.push("/login")

}
</script>

<template>
    <header>
        <nav class="flex flex-row space-between">
            <div>
                <RouterLink to="/">Home</RouterLink>
            </div>
            <div v-if="!userStore.user">
                <RouterLink  to="/login">Login</RouterLink>
                <RouterLink to="/register">Register</RouterLink>
            </div>
            <div v-else>
                <RouterLink v-if="userStore.user" to="/settings">{{ userStore.user.email }}</RouterLink>
                <a href="#" @click="logout">Logout</a>
            </div>
        </nav>
    </header>
</template>

<style scoped></style>

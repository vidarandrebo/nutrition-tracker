<script setup lang="ts">
import { useUserStore } from "../Stores/UserStore.ts";
import router from "../Router.ts";
import MenuBarRouterLink from "./MenuBarRouterLink.vue";

const userStore = useUserStore();

async function logout() {
    userStore.user = null;

    localStorage.removeItem("user");

    await router.push("/login");
}
</script>

<template>
    <header>
        <nav class="flex flex-row space-between">
            <div>
                <MenuBarRouterLink to="/">Home</MenuBarRouterLink>
                <MenuBarRouterLink to="/food-items">Food Items</MenuBarRouterLink>
                <MenuBarRouterLink to="/recipes">Recipes</MenuBarRouterLink>
            </div>
            <div v-if="!userStore.user">
                <MenuBarRouterLink to="/login">Login</MenuBarRouterLink>
                <MenuBarRouterLink to="/register">Register</MenuBarRouterLink>
            </div>
            <div v-else>
                <MenuBarRouterLink v-if="userStore.user" to="/settings">{{ userStore.user.email }}</MenuBarRouterLink>
                <a href="#" @click="logout">Logout</a>
            </div>
        </nav>
    </header>
</template>

<style scoped></style>

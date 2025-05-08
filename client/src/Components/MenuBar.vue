<script setup lang="ts">
import { useUserStore } from "../Stores/UserStore.ts";
import router from "../Router.ts";
import MenuBarRouterLink from "./MenuBarRouterLink.vue";
import { computed, ref } from "vue";

const userStore = useUserStore();

const burgerIsOpen = ref<boolean>(false);

const navbarBurgerClass = computed(() => {
    if (burgerIsOpen.value) {
        return "navbar-burger js-burger is-active";
    }
    return "navbar-burger js-burger";
});
const navbarMenuClass = computed(() => {
    if (burgerIsOpen.value) {
        return "navbar-menu is-active";
    }
    return "navbar-menu";
});

function flipBurger() {
    burgerIsOpen.value = !burgerIsOpen.value;
}

async function logout() {
    userStore.user = null;

    localStorage.removeItem("user");

    await router.push("/login");
}
</script>

<template>
    <header>
        <nav class="navbar is-transparent is-active">
            <div class="navbar-brand">
                <div :class="navbarBurgerClass" data-target="navbarExampleTransparentExample" @click="flipBurger">
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>

            <div id="navbarExampleTransparentExample" :class="navbarMenuClass">
                <div class="navbar-start">
                    <MenuBarRouterLink to="/">Home</MenuBarRouterLink>
                    <MenuBarRouterLink to="/food-items">Food Items</MenuBarRouterLink>
                    <MenuBarRouterLink to="/recipes">Recipes</MenuBarRouterLink>
                </div>

                <div class="navbar-end">
                    <MenuBarRouterLink v-if="!userStore.user" to="/login">Login</MenuBarRouterLink>
                    <MenuBarRouterLink v-if="!userStore.user" to="/register">Register</MenuBarRouterLink>
                    <MenuBarRouterLink v-if="userStore.user" to="/settings"
                        >{{ userStore.user.email }}
                    </MenuBarRouterLink>
                    <a href="#" @click="logout" class="button" v-if="userStore.user">Logout</a>
                </div>
            </div>
        </nav>
    </header>
</template>

<style scoped></style>

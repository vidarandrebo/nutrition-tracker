<script setup lang="ts">
import { useUserStore } from "../Stores/UserStore.ts";
import router from "../Router.ts";
import MenuBarRouterLink from "./MenuBarRouterLink.vue";
import { computed, ref } from "vue";
import { useMealStore } from "../Stores/MealStore.ts";
import { useFoodItemStore } from "../Stores/FoodItemStore.ts";

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
    const mealStore = useMealStore();
    mealStore.clear();
    const foodItemStore = useFoodItemStore();
    foodItemStore.clear();

    localStorage.removeItem("user");

    await router.push("/login");
}
</script>

<template>
    <header class="">
        <nav class="navbar is-transparent is-active has-background-primary-light">
            <div class="navbar-brand">
                <div :class="navbarBurgerClass" data-target="navbarExampleTransparentExample" @click="flipBurger">
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>

            <div id="navbarExampleTransparentExample" :class="navbarMenuClass" @click="flipBurger">
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
                    <a v-if="userStore.user" href="#" class="button" @click="logout">Logout</a>
                </div>
            </div>
        </nav>
    </header>
</template>

<style scoped></style>

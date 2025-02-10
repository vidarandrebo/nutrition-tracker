import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import LoginPage from "./pages/LoginPage.vue";
import RegisterPage from "./pages/RegisterPage.vue";
import SettingsPage from "./pages/SettingsPage.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: "Login",
        component: LoginPage,
    },
    {
        path: "/register",
        name: "Register",
        component: RegisterPage,
    },
    {
        path: "/settings",
        name: "Settings",
        component: SettingsPage,
    },
    {
        path: "/",
        name: "Home",
        component: HomePage,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});
export default router;

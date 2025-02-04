import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import LoginPage from "./pages/LoginPage.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: "Login",
        component: LoginPage,
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

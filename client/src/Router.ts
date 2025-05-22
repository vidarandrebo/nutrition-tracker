import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import HomePage from "./Pages/HomePage.vue";
import LoginPage from "./Pages/LoginPage.vue";
import RegisterPage from "./Pages/RegisterPage.vue";
import SettingsPage from "./Pages/SettingsPage.vue";
import RecipePage from "./Pages/RecipePage.vue";
import FoodItemPage from "./Pages/FoodItems/FoodItemPage.vue";
import AddFoodItemPage from "./Pages/FoodItems/AddFoodItemPage.vue";
import MealPage from "./Pages/Meals/MealPage.vue";

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
        path: "/food-items",
        name: "Food Items",
        component: FoodItemPage,
    },
    {
        path: "/food-items/add",
        name: "Add Food Item",
        component: AddFoodItemPage,
    },
    {
        path: "/recipes",
        name: "Recipes",
        component: RecipePage,
    },
    {
        path: "/",
        name: "Home",
        component: HomePage,
    },
    {
        path: "/meals/:id",
        name: "Meal",
        component: MealPage,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});
export default router;

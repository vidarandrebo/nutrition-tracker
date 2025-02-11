<script setup lang="ts">
import { reactive } from "vue";
import type { LoginForm } from "../Models/LoginForm.ts";
import InputText from "../Components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { AccessTokenResponse } from "../Models/AccessTokenResponse.ts";
import { type User, writeToLocalStorage } from "../Models/User.ts";
import router from "../Router.ts";
import { useUserStore } from "../Stores/UserStore.ts";

const userStore = useUserStore();
const loginForm = reactive<LoginForm>({ email: "", password: "" });

async function login() {
    const httpRequest = new HttpRequest()
        .setRoute("/api/login")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(loginForm);

    await httpRequest.send();
    const httpResponse = httpRequest.getResponseData();
    let loginResponse: AccessTokenResponse | undefined = undefined;

    if (httpResponse) {
        if (httpResponse?.status == 200) {
            loginResponse = httpResponse.body as AccessTokenResponse;
            const user: User = { email: loginForm.email, accessToken: loginResponse.token };
            writeToLocalStorage(user);
            userStore.user = user;
        }
    }
    await router.push("/");
}
</script>
<template>
    <h1>Login</h1>
    <form v-on:submit.prevent="login" class="flex flex-column">
        <label class="flex w-20 space-between pd-b-1">
            <p>Email</p>
            <InputText v-model="loginForm.email" type="email" />
        </label>
        <label class="flex w-20 space-between pd-b-1">
            <p>Password</p>
            <InputText v-model="loginForm.password" type="password" />
        </label>
        <button type="submit" class="w-10">Login</button>
    </form>
</template>

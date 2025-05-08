<script setup lang="ts">
import { reactive } from "vue";
import type { LoginForm } from "../Models/LoginForm.ts";
import InputText from "../Components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { AccessTokenResponse } from "../Models/AccessTokenResponse.ts";
import router from "../Router.ts";
import { useUserStore } from "../Stores/UserStore.ts";
import { User } from "../Models/User.ts";
import Button from "../Components/Button.vue";

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
            User.writeToLocalStorage(user);
            userStore.user = user;
        }
    }
    await router.push("/");
}
</script>
<template>
    <h1>Login</h1>
    <form v-on:submit.prevent="login">
        <label>
            <p>Email</p>
            <InputText v-model="loginForm.email" type="email" />
        </label>
        <label>
            <p>Password</p>
            <InputText v-model="loginForm.password" type="password" />
        </label>
        <button type="submit">Login</button>
        <Button type="submit">Login</Button>
    </form>
</template>

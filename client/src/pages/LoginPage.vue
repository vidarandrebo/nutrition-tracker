<script setup lang="ts">
import { reactive } from "vue";
import type { LoginForm } from "../models/LoginForm.ts";
import InputText from "../components/InputText.vue";
import { HttpRequest } from "http-methods-ts";

const loginForm = reactive<LoginForm>({ email: "", password: "" });

type AccessTokenResponse = {
    token: string;
}
async function login() {
    const httpRequest = new HttpRequest()
        .setRoute("/api/login")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(loginForm);

    await httpRequest.send();
    const httpResponse = httpRequest.getResponseData();
    let loginResponse: AccessTokenResponse | undefined = undefined

    if (httpResponse) {
        if (httpResponse?.status == 200) {
            loginResponse = httpResponse.body as AccessTokenResponse
        }
    }
    console.log(loginResponse?.token)
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

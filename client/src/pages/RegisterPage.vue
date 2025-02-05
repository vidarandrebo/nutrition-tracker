<script setup lang="ts">
import { reactive } from "vue";
import InputText from "../components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { RegisterForm } from "../models/RegisterForm.ts";

const registerForm = reactive<RegisterForm>({ email: "", password: "" });

async function register() {
    const httpRequest = new HttpRequest()
        .setRoute("/api/register")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setRequestData(registerForm);

    await httpRequest.send();
    const httpResponse = httpRequest.getResponseData();

    if (httpResponse) {
        if (httpResponse?.status == 201) {
            console.log("register successful");
        }
    }
}
</script>
<template>
    <h1>Register</h1>
    <form v-on:submit.prevent="register" class="flex flex-column">
        <label class="flex w-20 space-between pd-b-1">
            <p>Email</p>
            <InputText v-model="registerForm.email" type="email" />
        </label>
        <label class="flex w-20 space-between pd-b-1">
            <p>Password</p>
            <InputText v-model="registerForm.password" type="password" />
        </label>
        <button type="submit" class="w-10">Register</button>
    </form>
</template>

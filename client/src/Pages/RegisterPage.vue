<script setup lang="ts">
import { reactive } from "vue";
import InputText from "../Components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { RegisterForm } from "../Models/RegisterForm.ts";

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
    <form v-on:submit.prevent="register" >
        <label >
            <p>Email</p>
            <InputText v-model="registerForm.email" type="email" />
        </label>
        <label >
            <p>Password</p>
            <InputText v-model="registerForm.password" type="password" />
        </label>
        <button type="submit" >Register</button>
    </form>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import InputText from "../Components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { RegisterForm } from "../Models/RegisterForm.ts";
import HeaderH1 from "../Components/HeaderH1.vue";

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
    <HeaderH1>Register</HeaderH1>
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

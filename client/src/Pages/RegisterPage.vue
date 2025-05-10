<script setup lang="ts">
import { reactive } from "vue";
import InputText from "../Components/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { RegisterForm } from "../Models/RegisterForm.ts";
import HeaderH1 from "../Components/HeaderH1.vue";
import Button from "../Components/Button.vue";
import Label from "../Components/Label.vue";
import FormField from "../Components/FormField.vue";
import router from "../Router.ts";

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
            await router.push("/login");
        }
    }
}
</script>
<template>
    <HeaderH1>Register</HeaderH1>
    <div class="container ">
        <form v-on:submit.prevent="register" class="box">
            <FormField>
                <Label>
                    Email
                    <InputText v-model="registerForm.email" type="email" />
                </Label>
            </FormField>
            <FormField>
                <Label>
                    Password
                    <InputText v-model="registerForm.password" type="password" />
                </Label>
            </FormField>
            <Button type="submit">Register</Button>
        </form>
    </div>
</template>

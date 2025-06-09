<script setup lang="ts">
import { reactive } from "vue";
import type { LoginForm } from "../Models/LoginForm.ts";
import InputText from "../Components/Forms/InputText.vue";
import { HttpRequest } from "http-methods-ts";
import type { AccessTokenResponse } from "../Models/AccessTokenResponse.ts";
import router from "../Router.ts";
import { useUserStore } from "../Stores/UserStore.ts";
import { User } from "../Models/User.ts";
import ButtonPrimary from "../Components/Buttons/ButtonPrimary.vue";
import FormField from "../Components/Forms/FormField.vue";
import HeaderH1 from "../Components/HeaderH1.vue";
import Label from "../Components/Forms/Label.vue";

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
            await router.push("/");
        }
    }
}
</script>
<template>
    <HeaderH1>Login</HeaderH1>
    <div class="container ">
        <form v-on:submit.prevent="login" class="box">
            <FormField>
                <Label class="label">
                    Email
                    <div class="control">
                        <InputText v-model="loginForm.email" type="email" />
                    </div>
                </Label>
            </FormField>
            <FormField>
                <Label>
                    <p>Password</p>
                    <InputText v-model="loginForm.password" type="password" />
                </Label>
            </FormField>
            <FormField>
                <ButtonPrimary class="is-primary" type="submit">Login</ButtonPrimary>
            </FormField>
        </form>
    </div>
</template>

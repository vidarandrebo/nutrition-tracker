<script setup lang="ts">
import { reactive } from "vue";
import type { LoginForm } from "../Models/Auth/LoginForm.ts";
import InputText from "../Components/Forms/InputText.vue";
import router from "../Router.ts";
import { useUserStore } from "../Stores/UserStore.ts";
import { User } from "../Models/User.ts";
import ButtonPrimary from "../Components/Buttons/ButtonPrimary.vue";
import FormField from "../Components/Forms/FormField.vue";
import HeaderH1 from "../Components/Headings/HeaderH1.vue";
import LabelPrimary from "../Components/Forms/LabelPrimary.vue";
import { getAuthClient } from "../Models/Api.ts";

const userStore = useUserStore();
const loginForm = reactive<LoginForm>({ email: "", password: "" });

async function login() {
    const api = getAuthClient();
    const response = await api.apiLoginPost({
        loginRequest: {
            email: loginForm.email,
            password: loginForm.password,
        },
    });

    if (response.token) {
        const user: User = { email: loginForm.email, accessToken: response.token };
        User.writeToLocalStorage(user);
        userStore.user = user;
        await router.push("/");
    }
}
</script>
<template>
    <HeaderH1>Login</HeaderH1>
    <div class="container">
        <form class="box" @submit.prevent="login">
            <FormField>
                <LabelPrimary class="label">
                    Email
                    <div class="control">
                        <InputText v-model="loginForm.email" type="email" />
                    </div>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>Password</p>
                    <InputText v-model="loginForm.password" type="password" />
                </LabelPrimary>
            </FormField>
            <FormField>
                <ButtonPrimary class="is-primary" type="submit">Login</ButtonPrimary>
            </FormField>
        </form>
    </div>
</template>

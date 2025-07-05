<script setup lang="ts">
import { reactive } from "vue";
import InputText from "../Components/Forms/InputText.vue";
import type { RegisterForm } from "../Models/Auth/RegisterForm.ts";
import HeaderH1 from "../Components/Headings/HeaderH1.vue";
import ButtonPrimary from "../Components/Buttons/ButtonPrimary.vue";
import LabelPrimary from "../Components/Forms/LabelPrimary.vue";
import FormField from "../Components/Forms/FormField.vue";
import router from "../Router.ts";
import { getAuthClient } from "../Models/Api.ts";

const registerForm = reactive<RegisterForm>({ email: "", password: "" });

async function register() {
    const api = getAuthClient();
    try {
        await api.apiRegisterPost({ registerRequest: registerForm });
        console.log("register successful");
        await router.push("/login");
    } catch {
        console.log("failed to register");
    }
}
</script>
<template>
    <HeaderH1>Register</HeaderH1>
    <div class="container">
        <form class="box" @submit.prevent="register">
            <FormField>
                <LabelPrimary>
                    Email
                    <InputText v-model="registerForm.email" type="email" />
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    Password
                    <InputText v-model="registerForm.password" type="password" />
                </LabelPrimary>
            </FormField>
            <ButtonPrimary type="submit">Register</ButtonPrimary>
        </form>
    </div>
</template>

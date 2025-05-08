<script setup lang="ts">
import HeaderH1 from "../Components/HeaderH1.vue";
import { useUserStore } from "../Stores/UserStore.ts";
import { ref, watch } from "vue";
import InputDate from "../Components/InputDate.vue";
import { HttpRequest } from "http-methods-ts";

const userStore = useUserStore();

const selectedDay = ref(new Date());

type PostMealRequest = {
    timeStamp: Date;
};
watch(selectedDay, (newValue, oldValue) => {
    console.log(`Day changed from ${oldValue} to ${newValue}`);
});

async function addMeal() {
    const user = userStore.user;
    if (user === null) {
        return
    }
    console.log("adding meal");
    console.log(`today: ${isToday(selectedDay.value)}`);

    const request: PostMealRequest = {
        timeStamp: mealTimeStamp(),
    };
    console.log(request);

    const httpRequest = new HttpRequest()
        .setRoute("/api/meals")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken)
        .setRequestData(request);

    await httpRequest.send();

    const response = httpRequest.getResponseData();
    switch (response?.status) {
        case 201:
            console.log(response?.body)
            break;
        case 404:
            console.log("oi, ya goofed up")
            break;
        case 409:
        case 403:
            console.log("oida");
            break;
        default:
            break;
    }
}

function mealTimeStamp(): Date {
    const ts = isToday(selectedDay.value) ? new Date() : new Date(selectedDay.value);
    if (isNaN(ts.getUTCSeconds())) {
        return new Date();
    }
    return ts;
}

function isToday(date: Date): boolean {
    const now = new Date();
    return (
        date.getMonth() == now.getMonth() && date.getDate() == now.getDate() && date.getFullYear() == now.getFullYear()
    );
}
</script>
<template>
    <HeaderH1>Home</HeaderH1>
    <div v-if="userStore.user">
        <p>Welcome back {{ userStore.user.email }}</p>
        <InputDate v-model="selectedDay"></InputDate>
    </div>
    <button v-on:click="addMeal">Add meal</button>
</template>

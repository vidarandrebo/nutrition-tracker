<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { computed, reactive } from "vue";
import { PostFoodItemRequest } from "../../Models/FoodItems/Requests.ts";
import router from "../../Router.ts";
import { HttpRequest } from "http-methods-ts";
import { readFromLocalStorage } from "../../Models/User.ts";
import InputText from "../../Components/InputText.vue";
import InputNumber from "../../Components/InputNumber.vue";

const formModel = reactive<PostFoodItemRequest>(new PostFoodItemRequest());

async function postFoodItem() {
    const user = readFromLocalStorage()
    if (user === null) {
        return
    }
    const httpRequest = new HttpRequest()
        .setRoute("/api/food-items")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken)
        .setRequestData(formModel);

    await httpRequest.send();
    //const httpResponse = httpRequest.getResponseData();
    // if (httpResponse) {
    //        if (httpResponse?.status == 201) {
    //        }
    //}
    await router.push("/food-items");
}
const estKCal = computed(() => {
    return 4 * formModel.protein + 4 * formModel.carbohydrate + 9 * formModel.fat
})
</script>

<template>
    <HeaderH1>Add FoodItem</HeaderH1>
    <form v-on:submit.prevent="postFoodItem">
        <label>
            <p>Manufacturer</p>
            <InputText v-model="formModel.manufacturer"></InputText>
        </label>
        <label>
            <p>Product</p>
            <InputText v-model="formModel.product"></InputText>
        </label>
        <label>
            <p>Protein</p>
            <InputNumber v-model.number="formModel.protein"></InputNumber>
        </label>
        <label>
            <p>Carbohydrate</p>
            <InputNumber v-model.number="formModel.carbohydrate"></InputNumber>
        </label>
        <label>
            <p>Fat</p>
            <InputNumber v-model.number="formModel.fat"></InputNumber>
        </label>
        <label>
            <p>KCal</p>
            <InputNumber v-model.number="formModel.kCal" :place-holder="estKCal"></InputNumber>
        </label>
        <button type="submit">Add</button>
    </form>
</template>

<style scoped></style>

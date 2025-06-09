<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { computed, reactive } from "vue";
import { PostFoodItemRequest } from "../../Models/FoodItems/Requests.ts";
import router from "../../Router.ts";
import { HttpRequest } from "http-methods-ts";
import InputText from "../../Components/Forms/InputText.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import { FoodItem } from "../../Models/FoodItems/Fooditem.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import type { FoodItemResponse } from "../../Models/FoodItems/Responses.ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import FormField from "../../Components/Forms/FormField.vue";
import Label from "../../Components/Forms/Label.vue";

const formModel = reactive<PostFoodItemRequest>(new PostFoodItemRequest());

const userStore = useUserStore();
const foodItemStore = useFoodItemStore();

async function postFoodItem() {
    const user = userStore.user;
    if (user === null) {
        return;
    }
    const httpRequest = new HttpRequest()
        .setRoute("/api/food-items")
        .setMethod("POST")
        .addHeader("Content-Type", "application/json")
        .setBearerToken(user.accessToken)
        .setRequestData(formModel);

    await httpRequest.send();
    const httpResponse = httpRequest.getResponseData();
    if (httpResponse) {
        if (httpResponse?.status == 201) {
            const foodItem = FoodItem.fromResponse(httpResponse.body as FoodItemResponse);
            foodItemStore.collection.push(foodItem);
        }
    }
    await router.push("/food-items");
}

const estKCal = computed(() => {
    return 4 * formModel.protein + 4 * formModel.carbohydrate + 9 * formModel.fat;
});
</script>

<template>
    <HeaderH1>Add FoodItem</HeaderH1>
    <form v-on:submit.prevent="postFoodItem">
        <FormField>
            <Label>
                <p>Manufacturer</p>
                <InputText v-model="formModel.manufacturer"></InputText>
            </Label>
        </FormField>
        <FormField>
            <Label>
                <p>Product</p>
                <InputText v-model="formModel.product"></InputText>
            </Label>
        </FormField>
        <FormField>
            <Label>
                <p>Protein</p>
                <InputNumber v-model.number="formModel.protein"></InputNumber>
            </Label>
        </FormField>
        <FormField>
            <Label>
                <p>Carbohydrate</p>
                <InputNumber v-model.number="formModel.carbohydrate"></InputNumber>
            </Label>
        </FormField>
        <FormField>
            <Label>
                <p>Fat</p>
                <InputNumber v-model.number="formModel.fat"></InputNumber>
            </Label>
        </FormField>
        <FormField>
            <Label>
                <p>KCal</p>
                <InputNumber v-model.number="formModel.kCal" :place-holder="estKCal"></InputNumber>
            </Label>
        </FormField>
        <ButtonPrimary type="submit">Add</ButtonPrimary>
    </form>
</template>

<style scoped></style>

<script setup lang="ts">
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { computed, reactive } from "vue";
import { PostFoodItemRequest } from "../../Models/FoodItems/Requests.ts";
import router from "../../Router.ts";
import InputText from "../../Components/Forms/InputText.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import { FoodItem } from "../../Models/FoodItems/Fooditem.ts";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import FormField from "../../Components/Forms/FormField.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import { getFoodItemsClient } from "../../Models/Api.ts";

const formModel = reactive<PostFoodItemRequest>(new PostFoodItemRequest());

const foodItemStore = useFoodItemStore();

async function postFoodItem() {
    const client = getFoodItemsClient();

    try {
        const response = await client.apiFoodItemsPost({ postFoodItemRequest: formModel });
        const foodItem = FoodItem.fromResponse(response);
        foodItemStore.collection.push(foodItem);
        await router.push("/food-items");
    } catch {
        console.log("failed to create new food item");
    }
}

const estKCal = computed(() => {
    return 4 * formModel.protein + 4 * formModel.carbohydrate + 9 * formModel.fat;
});
</script>

<template>
    <div class="container">
        <HeaderH1>Add FoodItem</HeaderH1>
        <form @submit.prevent="postFoodItem">
            <FormField>
                <LabelPrimary>
                    <p>Manufacturer</p>
                    <InputText v-model="formModel.manufacturer"></InputText>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>Product</p>
                    <InputText v-model="formModel.product"></InputText>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>Protein</p>
                    <InputNumber v-model.number="formModel.protein"></InputNumber>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>Carbohydrate</p>
                    <InputNumber v-model.number="formModel.carbohydrate"></InputNumber>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>Fat</p>
                    <InputNumber v-model.number="formModel.fat" ></InputNumber>
                </LabelPrimary>
            </FormField>
            <FormField>
                <LabelPrimary>
                    <p>KCal</p>
                    <InputNumber v-model.number="formModel.kCal" :place-holder="estKCal"></InputNumber>
                </LabelPrimary>
            </FormField>
            <LevelPrimary>
                <template #right>
                    <ButtonPrimary type="submit">Add</ButtonPrimary>
                </template>
            </LevelPrimary>
        </form>
    </div>
</template>

<style scoped></style>

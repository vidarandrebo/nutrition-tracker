<script setup lang="ts">
import { useRoute } from "vue-router";
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted, ref } from "vue";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import HeaderH2 from "../../Components/Headings/HeaderH2.vue";
import ButtonDanger from "../../Components/Buttons/ButtonDanger.vue";
import { type PostFoodItemPortion } from "../../Gen";
import AddFoodItemPortion from "./AddFoodItemPortion.vue";

const route = useRoute();
let foodItemId = 0;

if (!Array.isArray(route.params.id)) {
    foodItemId = parseInt(route.params.id);
}

const foodItemStore = useFoodItemStore();

const foodItem = ref<FoodItem | undefined>(undefined);

onMounted(async () => {
    let localFi = foodItemStore.getFoodItem(foodItemId);
    if (!localFi) {
        const remoteFi = await FoodItem.getById(foodItemId);
        if (remoteFi) {
            foodItemStore.collection.push(remoteFi);
        }
        localFi = foodItemStore.getFoodItem(foodItemId);
    }
    foodItem.value = localFi;
});
const addPortionModalOpen = ref<boolean>(false);
async function addFoodItemPortion(ps: PostFoodItemPortion) {
    if (foodItem.value) {
        await foodItemStore.addPortionSize(foodItem.value.id, ps);
        addPortionModalOpen.value = false;
    }
}
</script>

<template>
    <section class="container">
        <template v-if="foodItem">
            <div class="is-flex is-justify-content-space-between">
                <HeaderH1>{{ foodItem.name }}</HeaderH1>
                <div class="field is-grouped">
                    <p class="control">
                        <ButtonPrimary>Edit</ButtonPrimary>
                    </p>
                    <p class="control">
                        <ButtonDanger>Delete</ButtonDanger>
                    </p>
                </div>
            </div>
            <div class="container">
                <HeaderH2>Nutrients</HeaderH2>
                <p>Protein: {{ foodItem.protein }}</p>
                <p>Carbohydrate: {{ foodItem.carbohydrate }}</p>
                <p>Fat: {{ foodItem.fat }}</p>
                <p>KCal: {{ foodItem.kCal }}</p>
            </div>
            <div class="container">
                <LevelPrimary>
                    <template #left>
                        <HeaderH2>Portions</HeaderH2>
                    </template>
                    <template #right>
                        <ButtonPrimary v-if="!addPortionModalOpen" @click="() => (addPortionModalOpen = true)"
                            >Add Portion</ButtonPrimary
                        >
                    </template>
                </LevelPrimary>
                <div v-for="portion in foodItem.portionSizes" :key="portion.id">
                    <p>{{ portion.name }} - {{ portion.amount }}g</p>
                </div>
                <AddFoodItemPortion
                    v-model="addPortionModalOpen"
                    @add-portion-size="addFoodItemPortion"
                ></AddFoodItemPortion>
            </div>
        </template>
    </section>
</template>

<style scoped></style>

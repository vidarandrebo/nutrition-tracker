<script setup lang="ts">
import { useRoute } from "vue-router";
import HeaderH1 from "../../Components/Headings/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted } from "vue";
import { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import { ref } from "vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";

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
</script>

<template>
    <section class="container">
        <template v-if="foodItem">
            <LevelPrimary>
                <template #left>
                    <HeaderH1>{{ foodItem.name }}</HeaderH1>
                </template>
                <template #right>
                    <ButtonPrimary>Edit</ButtonPrimary>
                    <ButtonPrimary>Delete</ButtonPrimary>
                </template>
            </LevelPrimary>
            <div class="container">
                <p>Protein</p>
                <p>{{ foodItem.protein }}</p>
            </div>
        </template>
    </section>
</template>

<style scoped></style>

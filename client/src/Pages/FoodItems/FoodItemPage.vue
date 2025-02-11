<script setup lang="ts">
import HeaderH1 from "../../Components/HeaderH1.vue";
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import { onMounted } from "vue";

const foodItemStore = useFoodItemStore();

onMounted(async () => {
    await foodItemStore.init();
});
</script>

<template>
    <HeaderH1>Food Items</HeaderH1>
    <table>
        <thead>
            <tr>
                <th>Id</th>
                <th>Manufacturer</th>
                <th>Product</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in foodItemStore.collection" :key="item.id">
                <td>{{ item.id }}</td>
                <td>{{ item.manufacturer }}</td>
                <td>{{ item.product }}</td>
            </tr>
        </tbody>
    </table>
</template>

<style scoped></style>

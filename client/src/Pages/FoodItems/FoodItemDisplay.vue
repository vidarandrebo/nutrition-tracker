<script setup lang="ts">


import type { FoodItem } from "../../Models/FoodItems/Fooditem.ts";

type FoodItemDisplayProps = {
    item: FoodItem;
}
const props = defineProps<FoodItemDisplayProps>();
</script>

<template>
    <li class="box">
        <p>Manufacturer: {{ props.item.manufacturer}}</p>
        <p>Product {{ props.item.product }}</p>
        <p>Protein {{ props.item.protein }}</p>
        <p>Carbohydrate {{ props.item.carbohydrate }}</p>
        <p>Fat {{ props.item.fat }}</p>
        <p>KCal {{ props.item.kCal }}</p>
    </li>
</template>

<style scoped>

</style>
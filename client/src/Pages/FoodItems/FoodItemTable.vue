<script setup lang="ts">
import { BTable, BTableColumn } from "buefy";
import type { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import { ref } from "vue";

defineProps<{
    foodItems: FoodItem[];
}>();

const paginationPosition = "both";
const sortIcon = ref<string>("menu-up");
const currentPage = ref<number>(1);
const perPage = 20;

const checked = ref<FoodItem[]>([]);

function flipSortCaret(_field?: string, order?: string) {
    if (order === "asc") {
        sortIcon.value = "menu-up";
    }
    if (order === "desc") {
        sortIcon.value = "menu-down";
    }
}
</script>

<template>
    <BTable
        v-model:current-page="currentPage"
        v-model:checked-rows.sync="checked"
        hoverable
        checkable
        :data="foodItems"
        paginated
        :per-page="perPage"
        :pagination-position="paginationPosition"
        :sort-icon="sortIcon"
        default-sort="name"
        aria-next-label="Next page"
        aria-previous-label="Previous page"
        aria-page-label="Page"
        aria-current-label="Current page"
        @sort="flipSortCaret"
    >
        <BTableColumn v-slot="slotProps" label="Id" sortable field="id" width="40">
            {{ slotProps.row.id }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" label="Name" sortable field="name">
            {{ slotProps.row.name }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" sortable field="protein" label="Protein">
            {{ slotProps.row.protein }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" sortable field="carbohydrate" label="Carbohydrate">
            {{ slotProps.row.carbohydrate }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" sortable field="fat" label="Fat">
            {{ slotProps.row.fat }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" sortable field="kCal" label="KCal">
            {{ slotProps.row.kCal }}
        </BTableColumn>
    </BTable>
</template>

<style>
table {
}
</style>

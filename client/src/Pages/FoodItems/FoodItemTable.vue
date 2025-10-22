<script setup lang="ts">
import { BButton, BField, BInput, BSelect, BSwitch, BTable, BTableColumn } from "buefy";
import type { FoodItem } from "../../Models/FoodItems/FoodItem.ts";
import { ref } from "vue";

const props = defineProps<{
    foodItems: FoodItem[];
}>();

const isPaginated = ref<boolean>(true);
const paginationPosition = ref<string>("top");
const isPaginationSimple = ref<boolean>(false);
const isPaginationRounded = ref<boolean>(false);
const defaultSortDirection = ref<string>("asc");
const sortIcon = ref<string>("arrow-up");
const sortIconSize = ref<string>("is-small");
const currentPage = ref<number>(1);
const perPage = ref<number>(20);
const hasInput = ref<boolean>(false);
const paginationOrder = ref("");
const inputPosition = ref("");
const inputDebounce = ref("");
</script>

<template>
    <b-field grouped group-multiline>
        <b-select v-model="defaultSortDirection">
            <option value="asc">Default sort direction: ASC</option>
            <option value="desc">Default sort direction: DESC</option>
        </b-select>
        <b-select v-model="perPage" :disabled="!isPaginated">
            <option value="5">5 per page</option>
            <option value="10">10 per page</option>
            <option value="15">15 per page</option>
            <option value="20">20 per page</option>
        </b-select>
        <div class="control">
            <b-button label="Set page to 2" :disabled="!isPaginated" @click="currentPage = 2" />
        </div>
        <div class="control is-flex">
            <b-switch v-model="isPaginated">Paginated</b-switch>
        </div>
        <div class="control is-flex">
            <b-switch v-model="isPaginationSimple" :disabled="!isPaginated">Simple pagination</b-switch>
        </div>
        <div class="control is-flex">
            <b-switch v-model="isPaginationRounded" :disabled="!isPaginated">Rounded pagination</b-switch>
        </div>
        <b-select v-model="paginationPosition" :disabled="!isPaginated">
            <option value="bottom">bottom pagination</option>
            <option value="top">top pagination</option>
            <option value="both">both</option>
        </b-select>
        <b-select v-model="sortIcon">
            <option value="arrow-up">Arrow sort icon</option>
            <option value="menu-up">Caret sort icon</option>
            <option value="chevron-up">Chevron sort icon</option>
        </b-select>
        <b-select v-model="sortIconSize">
            <option value="is-small">Small sort icon</option>
            <option value="">Regular sort icon</option>
            <option value="is-medium">Medium sort icon</option>
            <option value="is-large">Large sort icon</option>
        </b-select>
        <b-select v-model="paginationOrder">
            <option value="">default pagination order</option>
            <option value="is-centered">is-centered pagination order</option>
            <option value="is-right">is-right pagination order</option>
        </b-select>
        <div class="control is-flex">
            <b-switch v-model="hasInput">Input</b-switch>
        </div>
        <b-select v-model="inputPosition">
            <option value="">default input position</option>
            <option value="is-input-right">is-input-right</option>
            <option value="is-input-left">is-input-left</option>
        </b-select>
        <b-input v-model="inputDebounce" type="number" placeholder="debounce (milliseconds)" min="0"></b-input>
    </b-field>
    <BTable
        v-model:current-page="currentPage"
        :data="foodItems"
        :paginated="isPaginated"
        :per-page="perPage"
        :pagination-simple="isPaginationSimple"
        :pagination-position="paginationPosition"
        :pagination-rounded="isPaginationRounded"
        :sort-icon="sortIcon"
        :sort-icon-size="sortIconSize"
        default-sort="id"
        aria-next-label="Next page"
        aria-previous-label="Previous page"
        aria-page-label="Page"
        aria-current-label="Current page"
        :page-input="hasInput"
        :pagination-order="paginationOrder"
        :page-input-position="inputPosition"
        :debounce-page-input="inputDebounce"
    >
        <BTableColumn v-slot="slotProps" label="Id" sortable field="id">
            {{ slotProps.row.id }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" label="Name" sortable field="name">
            {{ slotProps.row.name }}
        </BTableColumn>
        <BTableColumn
            v-slot="slotProps"
            sortable
            field="protein"
            cell-class="is-hidden-mobile"
            header-class="is-hidden-mobile"
            label="Protein"
        >
            {{ slotProps.row.protein }}
        </BTableColumn>
        <BTableColumn
            v-slot="slotProps"
            sortable
            field="carbohydrate"
            cell-class="is-hidden-mobile"
            header-class="is-hidden-mobile"
            label="Carbohydrate"
        >
            {{ slotProps.row.carbohydrate }}
        </BTableColumn>
        <BTableColumn
            v-slot="slotProps"
            sortable
            field="fat"
            cell-class="is-hidden-mobile"
            header-class="is-hidden-mobile"
            label="Fat"
        >
            {{ slotProps.row.fat }}
        </BTableColumn>
        <BTableColumn v-slot="slotProps" sortable field="kCal" label="KCal">
            {{ slotProps.row.kCal }}
        </BTableColumn>
    </BTable>
</template>

<style>
table {
    width: 100%;
}
</style>

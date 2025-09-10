<script setup lang="ts">
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import InputText from "../../Components/Forms/InputText.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import { onMounted } from "vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import ButtonPlain from "../../Components/Buttons/ButtonPlain.vue";
import HeaderH2 from "../../Components/Headings/HeaderH2.vue";
import { useFilterStore } from "../../Stores/FilterStore.ts";
import FormField from "../../Components/Forms/FormField.vue";

const foodItemStore = useFoodItemStore();
const filterStore = useFilterStore();

const emit = defineEmits<{
    select: [id: number];
    cancel: [];
}>();

function cancel() {
    emit("cancel");
}

onMounted(async () => {
    await foodItemStore.init();
});
</script>

<template>
    <section>
        <LevelPrimary>
            <template #left>
                <HeaderH2>Add Food item</HeaderH2>
            </template>
            <template #right>
                <ButtonPlain class="level-item" @click="cancel">Cancel</ButtonPlain>
            </template>
        </LevelPrimary>
        <div class="is-flex is-align-items-center is-gap-1">
            <FormField class="is-flex-grow-1">
                <LabelPrimary>
                    Search (showing {{ foodItemStore.filteredFoodItems.length }} entries)
                    <InputText v-model="filterStore.foodItem.searchTerm" placeholder="Search"></InputText>
                </LabelPrimary>
            </FormField>
            <FormField>
                <label class="checkbox">
                    Show public
                    <input v-model="filterStore.foodItem.showPublic" type="checkbox" />
                </label>
            </FormField>
        </div>
        <div
            v-for="item in foodItemStore.filteredFoodItems"
            :key="item.id"
            class="is-flex is-flex-direction-row is-justify-content-space-between box"
        >
            <p>
                <b>{{ item.name }}</b>
            </p>
            <p class="pr-2"><b>KCal:</b> {{ item.kCal }}</p>
            <ButtonPrimary @click="() => emit('select', item.id)">Add</ButtonPrimary>
        </div>
    </section>
</template>

<style scoped></style>

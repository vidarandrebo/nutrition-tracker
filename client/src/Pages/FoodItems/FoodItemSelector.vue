<script setup lang="ts">
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import InputText from "../../Components/Forms/InputText.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import { onMounted } from "vue";
import LevelPrimary from "../../Components/LevelPrimary.vue";
import ButtonPlain from "../../Components/Buttons/ButtonPlain.vue";
import HeaderH2 from "../../Components/HeaderH2.vue";

const foodItemStore = useFoodItemStore();

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
    <section class="section">
        <LevelPrimary>
            <template #left>
                <HeaderH2>Add Food item</HeaderH2>
            </template>
            <template #right>
                <ButtonPlain class="level-item" @click="cancel">Cancel</ButtonPlain>
            </template>
        </LevelPrimary>
        <LabelPrimary>
            Search (showing {{ foodItemStore.filteredFoodItems.length }} entries)
            <InputText v-model="foodItemStore.searchTerm" placeholder="Search"></InputText>
        </LabelPrimary>
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

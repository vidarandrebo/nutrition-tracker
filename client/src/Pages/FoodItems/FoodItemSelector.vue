<script setup lang="ts">
import { useFoodItemStore } from "../../Stores/FoodItemStore.ts";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import InputText from "../../Components/Forms/InputText.vue";
import Label from "../../Components/Forms/Label.vue";
import { onMounted } from "vue";
import Level from "../../Components/Level.vue";
import Button from "../../Components/Buttons/Button.vue";
import HeaderH2 from "../../Components/HeaderH2.vue";

const foodItemStore = useFoodItemStore();

const emit = defineEmits<{
    select: [id: number]
    cancel:[]
}>();

function cancel() {
    emit('cancel')
}
onMounted(async () => {
    await foodItemStore.init();
});
</script>

<template>
    <section class="section">
        <Level>
            <template #left>
                <HeaderH2>Add Food item</HeaderH2>
            </template>
            <template #right>
                <Button @click="cancel" class="level-item">Cancel</Button>
            </template>
        </Level>
        <Label >
            Search (showing {{ foodItemStore.filteredFoodItems.length }} entries)
            <InputText v-model="foodItemStore.searchTerm" placeholder="Search"></InputText>
        </Label>
        <div v-for="item in foodItemStore.filteredFoodItems"
             :key="item.id"
             class="is-flex is-flex-direction-row is-justify-content-space-between box">
            <p><b>{{ item.name }}</b></p>
            <p class="pr-2"><b>KCal:</b> {{ item.kCal }}</p>
            <ButtonPrimary @click="() => emit('select', item.id)">Add</ButtonPrimary>
        </div>
    </section>
</template>

<style scoped>

</style>
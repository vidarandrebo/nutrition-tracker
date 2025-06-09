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
        <HeaderH2>Add Food item</HeaderH2>
        <Level>
            <template #left>
                <Label>
                    Search (showing {{ foodItemStore.filteredFoodItems.length }} entries)
                    <InputText v-model="foodItemStore.searchTerm"></InputText>
                </Label>
            </template>
            <template #right>
                <Button @click="cancel">Cancel</Button>
            </template>
        </Level>
        <div v-for="item in foodItemStore.filteredFoodItems"
             class="is-flex is-flex-direction-row is-justify-content-space-between box">
            <p>{{ item.name }}</p>
            <ButtonPrimary @click="() => emit('select', item.id)">Add</ButtonPrimary>
        </div>
    </section>
</template>

<style scoped>

</style>
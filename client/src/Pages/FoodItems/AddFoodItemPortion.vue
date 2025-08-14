<script setup lang="ts">
import InputText from "../../Components/Forms/InputText.vue";
import ModalPrimary from "../../Components/ModalPrimary.vue";
import LabelPrimary from "../../Components/Forms/LabelPrimary.vue";
import ButtonPrimary from "../../Components/Buttons/ButtonPrimary.vue";
import InputNumber from "../../Components/Forms/InputNumber.vue";
import { ref, watch } from "vue";
import type { PortionSizeForm } from "../../Models/FoodItems/PortionSize.ts";

const emit = defineEmits<{
    addPortionSize: [entry: PortionSizeForm];
}>();

const model = defineModel<boolean>({ required: true });
const addPortionModalOpen = ref<boolean>(model.value);
const formData = ref<PortionSizeForm>({
    amount: 0,
    name: "",
});

watch(addPortionModalOpen, (v) => {
    model.value = v;
});
watch(model, (v) => {
    addPortionModalOpen.value = v;
});
function submit() {
    emit("addPortionSize", formData.value);
}
</script>

<template>
    <ModalPrimary v-model="addPortionModalOpen" title="Add portionsize">
        <template #default>
            <div class="is-flex is-flex-direction-column">
                <LabelPrimary>
                    <p>Name</p>
                    <InputText v-model="formData.name"></InputText>
                </LabelPrimary>
                <LabelPrimary>
                    <p>Amount</p>
                    <InputNumber v-model="formData.amount"></InputNumber>
                </LabelPrimary>
            </div>
        </template>
        <template #footer>
            <ButtonPrimary @click="submit">Add</ButtonPrimary>
        </template>
    </ModalPrimary>
</template>

<style scoped></style>

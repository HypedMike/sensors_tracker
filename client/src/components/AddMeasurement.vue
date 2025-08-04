<template>
    <div class="mt-4 mb-4">
        <label for="measurement" class="block mb-2">New Measurement Value:</label>
        <div style="display: grid; grid-template-columns: 1fr auto; gap: 10px;" class="w-full">
            <input
                v-model="newMeasurement.value"
                type="number"
                placeholder="Enter measurement value"
                class="border p-2 rounded w-full"
                name="measurement"
            />
            <button
                @click="addMeasurement"
                class="flex items-center justify-center text-white p-2 rounded transition-colors duration-200 cursor-pointer"
            >
                <img src="/icons/plus.png" alt="Add Measurement" class="inline-block h-10 w-10 mr-2" />
            </button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const newMeasurement = ref({
    value: null
});

const emit = defineEmits(['addMeasurement']);

function addMeasurement() {
    if (newMeasurement.value.value !== null && newMeasurement.value.value !== undefined) {
        // Emit the new measurement value to the parent component
        emit('addMeasurement', newMeasurement.value);
        // Reset the input field
        newMeasurement.value = { value: null };
    } else {
        alert('Please enter a valid measurement value.');
    }
}

</script>
<script setup lang="ts">
import { ref } from 'vue';
import Sensor from '../models/sensor';

let sensor = ref<Sensor>(new Sensor({}));
let address = ref<{
    street: string;
    city: string;
    state: string;
    zip: string;
}>({
    street: '',
    city: '',
    state: '',
    zip: ''
});

const types = [
    { value: 'temperature', label: 'Temperature' },
    { value: 'humidity', label: 'Humidity' },
    { value: 'pressure', label: 'Pressure' }
];

function createSensor() {
    // Here you would typically send the sensor data to your API
    console.log('Creating sensor:', sensor.value, address.value);
    // Reset form after submission
    sensor.value = new Sensor({});
    address.value = { street: '', city: '', state: '', zip: '' };
}


</script>

<template>
    <div class="p-4">
        <h1 class="text-2xl font-bold mb-4">{{ $t('sensors.create.title') }}</h1>
        <form @submit.prevent="createSensor">
            <div class="mb-4">
                <label for="id" class="block text-sm font-medium text-gray-700">Id</label>
                <input v-model="sensor.id" type="text" id="id" 
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                focus:border-indigo-500 focus:ring-indigo-500" required>
                <span class="text-xs text-gray-500">{{ $t('sensors.create.id_left_empty') }}</span>
            </div>
            <div class="mb-4">
                <label for="name" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.name') }}</label>
                <input v-model="sensor.name" type="text" id="name" 
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                focus:border-indigo-500 focus:ring-indigo-500" required>
            </div>
            <div class="mb-4">
                <label for="type" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.type') }}</label>
                <select v-model="sensor.type" id="type" class="mt-1 block w-full border-gray-300 p-2
                rounded-md shadow-sm focus:border-indigo-500 focus:ring-indigo-500" required>
                    <option v-for="type in types" :key="type.value" :value="type.value">{{ type.label }}</option>
                </select>
            </div>
            <div class="mb-4">
                <label for="location" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.location') }}</label>
                <div class="grid grid-cols-2 gap-4">
                    <input v-model="address.street" type="text" placeholder="Street"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                    focus:border-indigo-500 focus:ring-indigo-500 mb-2" required>
                    <input v-model="address.city" type="text" placeholder="City"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                    focus:border-indigo-500 focus:ring-indigo-500 mb-2" required>
                    <input v-model="address.state" type="text" placeholder="State"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                    focus:border-indigo-500 focus:ring-indigo-500" required>
                    <input v-model="address.zip" type="number" placeholder="Zip Code"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                    focus:border-indigo-500 focus:ring-indigo-500" required>
                </div>
            </div>
            <button type="submit" 
            style="background-color: var(--secondary);"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
                {{ $t('sensors.create.crt_button') }}
            </button>
        </form>
    </div>
</template>
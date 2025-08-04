<script setup lang="ts">
import { ref } from 'vue';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import { useRouter } from 'vue-router';

let sensor = ref<Sensor>(new Sensor({}));

const route = useRouter();

function createSensor() {
    sensorsApi.createSensor(sensor.value).then(() => {
        // Handle successful creation
        alert('Sensor created successfully!');
        // Remove the sensor from the form
        sensor.value = new Sensor({});
        // Redirect to the details page
        route.push(`/sensors/${sensor.value.id}`);
    }).catch(error => {
        console.error('Error creating sensor:', error);
        alert('Failed to create sensor. Please try again.');
    });
}


</script>

<template>
    <div class="">
        <h1 class="text-2xl font-bold mb-4">{{ $t('sensors.create.title') }}</h1>
        <form @submit.prevent="createSensor">
            <div class="mb-4">
                <label for="name" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.name') }}</label>
                <input v-model="sensor.name" type="text" id="name" 
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                focus:border-indigo-500 focus:ring-indigo-500" required>
            </div>
            <div class="mb-4">
                <label for="threshold" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.threshold') }}</label>
                <input v-model.number="sensor.threshold" type="number" id="threshold"
                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                focus:border-indigo-500 focus:ring-indigo-500" required>
            </div>
            <div class="mb-4">
                <table class="w-full">
                    <tbody>
                        <tr class="w-full">
                            <td class="w-1/3">
                                <label for="latitude" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.latitude') }}</label>
                            </td>
                            <td class="w-full">
                                <input v-model.number="sensor.location[0]" type="number" id="latitude" placeholder="Latitude" min="-90" max="90"
                                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                                focus:border-indigo-500 focus:ring-indigo-500" required>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                <label for="longitude" class="block text-sm font-medium text-gray-700">{{ $t('sensors.create.longitude') }}</label>
                            </td>
                            <td>
                                <input v-model.number="sensor.location[1]" type="number" id="longitude" placeholder="Longitude" min="-180" max="180"
                                class="mt-1 block w-full border-gray-300 rounded-md shadow-sm p-2
                                focus:border-indigo-500 focus:ring-indigo-500" required>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <button type="submit" 
            style="background-color: var(--secondary);"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
                {{ $t('sensors.create.crt_button') }}
            </button>
        </form>
    </div>
</template>
<template>
    <div>
        <h1 class="text-2xl font-bold mb-4">{{ $t('sensors.list.title') }}</h1>
        <div class="grid grid-cols-5 gap-4 mb-4">
            <input type="text" v-model="query" placeholder="Search sensors..." class="border border-gray-300 rounded-lg p-2 mb-4 w-full h-12 col-span-4" />
            <button type="button" class="bg-blue-500 text-white rounded-lg h-12 col-span-1" value="Search">Search</button>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <SensorCard v-for="sensor in filteredSensors" :key="sensor.id" :sensor="sensor" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import SensorCard from '../components/SensorCard.vue';

const query = ref('');
const sensors = ref<Sensor[]>([]);
const filteredSensors = computed(() => {
    if (!query.value) return sensors.value;
    return sensors.value.filter(sensor => sensor.name.toLowerCase().includes(query.value.toLowerCase()));
})

onMounted(() => {
    sensorsApi.getSensors(true).then(response => {
        sensors.value = response;
    }).catch(error => {
        console.error('Error fetching sensors:', error);
    });
})
</script>
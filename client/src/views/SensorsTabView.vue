<template>
    <div>
        <h1 class="text-2xl font-bold mb-4">{{ $t('sensors.list.title') }}</h1>
        <div class="grid grid-cols-5 gap-4 mb-4">
            <input type="text" v-model="query" placeholder="Search sensors..." class="border border-gray-300 rounded-lg p-2 mb-4 w-full h-12 col-span-4" />
            <button type="button" class="bg-blue-500 text-white rounded-lg h-12 col-span-1" value="Search">Search</button>
        </div>
        <div v-if="loading" class="text-center py-4 h-screen flex items-center justify-center">
            <VueSpinnerBall type="TailSpin" color="#000" size="40" />
        </div>
        <div v-else>
            <table class="min-w-full bg-white border border-gray-300">
                <thead class="bg-gray-200">
                    <tr class="text-left">
                        <th class="px-4 py-2 cursor-pointer" @click="sortBy('id')">
                            {{ $t('sensors.list.id') }}
                            <span v-if="sortKey === 'id'" class="ml-2">
                                <svg :xmlns="plusIcon" class="h-4 w-4 inline" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width={2} d="M12 4v16m8-8H4" />
                                </svg>
                            </span>
                        </th>
                        <th class="px-4 py-2 cursor-pointer" @click="sortBy('name')">
                            {{ $t('sensors.list.name') }}
                            <span v-if="sortKey === 'name'" class="ml-2">
                                <svg :xmlns="plusIcon" class="h-4 w-4 inline" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width={2} d="M12 4v16m8-8H4" />
                                </svg>
                            </span>
                        </th>
                        <th class="px-4 py-2 cursor-pointer" @click="sortBy('latest')">
                            {{ $t('sensors.list.last_measurement') }}
                            <span v-if="sortKey === 'latest'" class="ml-2">
                                <svg :xmlns="plusIcon" class="h-4 w-4 inline" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width={2} d="M12 4v16m8-8H4" />
                                </svg>
                            </span>
                        </th>
                        <th class="px-4 py-2 cursor-pointer" @click="sortBy('status')">
                            {{ $t('sensors.list.status') }}
                            <span v-if="sortKey === 'status'" class="ml-2">
                                <svg :xmlns="plusIcon" class="h-4 w-4 inline" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width={2} d="M12 4v16m8-8H4" />
                                </svg>
                            </span>
                        </th>
                    </tr>
                </thead>
                <tbody class="text-gray-700">
                    <tr v-for="sensor in filteredSensors" :key="sensor.id" 
                        @click="goToSensorDetail(sensor)"
                        :class="`hover:bg-gray-100 cursor-pointer ${sensor.status() === 'ALARM' ? 'bg-red-100' : ''}`">
                        <td class="border border-gray-300 px-4 py-2">{{ sensor.id }}</td>
                        <td class="border border-gray-300 px-4 py-2">{{ sensor.name }}</td>
                        <td class="border border-gray-300 px-4 py-2">
                            {{ sensor.measurements.length > 0 ? sensor.measurements[sensor.measurements.length - 1].value.toFixed(2) : 'No data' }}
                        </td>
                        <td class="border border-gray-300 px-4 py-2">{{ sensor.status() }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import { VueSpinnerBall } from 'vue3-spinners';
import { useRouter } from 'vue-router';

const query = ref('');
const sensors = ref<Sensor[]>([]);
const loading = ref<boolean>(true);

const sortKey = ref<'id' | 'name' | 'latest' | 'status'>('id');
const sortOrder = ref<'asc' | 'desc'>('asc');

const plusIcon = "http://www.w3.org/2000/svg";

function sortBy(key: 'id' | 'name' | 'latest' | 'status') {
    if (sortKey.value === key) {
        sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc';
    } else {
        sortKey.value = key;
        sortOrder.value = 'asc';
    }
}

const filteredSensors = computed(() => {
    let result = sensors.value;
    if (query.value) {
        result = result.filter(sensor => sensor.name.toLowerCase().includes(query.value.toLowerCase()));
    }
    // Sorting
    return [...result].sort((a, b) => {
        let aValue, bValue;
        switch (sortKey.value) {
            case 'id':
                aValue = a.id;
                bValue = b.id;
                break;
            case 'name':
                aValue = a.name.toLowerCase();
                bValue = b.name.toLowerCase();
                break;
            case 'latest':
                aValue = a.measurements.length > 0 ? a.measurements[a.measurements.length - 1].value : -Infinity;
                bValue = b.measurements.length > 0 ? b.measurements[b.measurements.length - 1].value : -Infinity;
                break;
            case 'status':
                aValue = a.status();
                bValue = b.status();
                break;
        }
        if (aValue < bValue) return sortOrder.value === 'asc' ? -1 : 1;
        if (aValue > bValue) return sortOrder.value === 'asc' ? 1 : -1;
        return 0;
    });
});

const router = useRouter();

onMounted(() => {
    sensorsApi.getSensors(true).then(response => {
        sensors.value = response;
    }).catch(error => {
        console.error('Error fetching sensors:', error);
    }).finally(() => {
        loading.value = false;
    });
})

function goToSensorDetail(sensor: Sensor) {
    router.push({ name: 'sensorDetails', params: { id: sensor.id } });
}
</script>
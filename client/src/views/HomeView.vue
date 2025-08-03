<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import { VueSpinnerBall } from 'vue3-spinners';
import { useUserStore } from '../stores/userStore';

const division = ref<HTMLElement | null>(null);
const sensors = ref<Sensor[]>([]);
let measurements = computed<{
    id: string;
    value: number;
    sensor: Sensor;
    timestamp: Date;
}[]>(() => {
    let result: {
        id: string;
        value: number;
        sensor: Sensor;
        timestamp: Date;
    }[] = [];
    sensors.value.forEach(sensor => {
        sensor.measurements.forEach(measurement => {
            result.push({
                id: measurement.id,
                value: measurement.value,
                sensor: sensor,
                timestamp: measurement.timestamp
            });
        });
    });

    return result
    .sort((a, b) => b.timestamp.getTime() - a.timestamp.getTime())
    .slice(0, 30);
});


let loading = ref(true);
let center = ref<[number, number]>([17.425106510223976, 42.05612196520227]);
let screenWidth = window.innerWidth;

onMounted(() => {
  sensorsApi.getSensors(true).then(response => {
    sensors.value = response;
  }).catch(error => {
    console.error('Error fetching sensors:', error);
  }).finally(() => {
    loading.value = false;
  });
})

function calculateColor(sensor: Sensor): string {
    const overThreshold = sensor.measurements.filter(m => m.value > sensor.threshold).length;
    const totalMeasurements = sensor.measurements.length;
    const percentage = totalMeasurements > 0 ? ((100 - overThreshold) / totalMeasurements) * 100 : 0;
    let color = 'gray'; // default color
    if (percentage > 90) {
        color = 'green';
    } else if (percentage > 80) {
        color = 'yellow';
    } else if (percentage > 25) {
        color = 'red';
    }
    return color;
}

const userStore = useUserStore();

</script>

<template>
  <h1 class="text-2xl font-bold mb-4">{{ $t('dashboard.hello', { name: userStore.user?.username }) }}</h1>
  <div class="h-full" ref="division" v-if="!loading">
    <ol-map :style="`min-width: ${Math.min(screenWidth - 40, 400)}px; height: 400px;`">
      <ol-view :center="center" :zoom="5" projection="EPSG:4326" />
      <ol-tile-layer>
        <ol-source-osm />
      </ol-tile-layer>
      <ol-vector-layer>
        <ol-source-vector>
          <ol-feature v-for="sensor in sensors" :key="sensor.id">
            <ol-geom-point :coordinates="sensor.location"></ol-geom-point>
            <ol-style>
              <ol-style-circle :radius="1">
                <ol-style-fill :color="calculateColor(sensor)"></ol-style-fill>
                <ol-style-stroke :color="calculateColor(sensor)" width="10"></ol-style-stroke>
              </ol-style-circle>
            </ol-style>
          </ol-feature>
        </ol-source-vector>
      </ol-vector-layer>
    </ol-map>

    <section class="pb-4 mt-4">
      <h2 class="text-xl font-semibold mt-4 mb-2">
        {{ $t('dashboard.latest_measurements') }}
      </h2>
      <table class="min-w-full bg-white shadow-md rounded-lg overflow-hidden p-2"
        :style="`width: ${Math.min(screenWidth - 40, 400)}px;`">
        <thead>
          <tr class="bg-gray-300">
            <th class="text-left px-2">{{ $t('sensors.list.sensor_name') }}</th>
            <th class="text-left px-2">{{ $t('sensors.list.value') }}</th>
            <th class="text-left px-2">{{ $t('sensors.list.timestamp') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="measurement in measurements" :key="measurement.id" class="hover:bg-gray-100 transition-colors duration-200">
            <td class="px-2">{{ measurement.sensor.name }}</td>
            <td class="px-2">{{ measurement.value.toFixed(2) }}</td>
            <td class="px-2">{{ new Date(measurement.timestamp).toLocaleString() }}</td>
          </tr>
        </tbody>
      </table>
    </section>
  </div>
  <div v-else class="flex justify-center items-center h-[80vh]">
    <VueSpinnerBall size="50" color="red" />
  </div>
</template>
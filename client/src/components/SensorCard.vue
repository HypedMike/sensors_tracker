<template>
    <div class="bg-white shadow p-4 border border-gray-200 rounded-lg grid grid-cols-1 gap-4 
    hover:shadow-lg transition-shadow duration-200 cursor-pointer" @click="() => $router.push({ name: 'sensorDetails', params: { id: sensor.id } })">
        <section>
            <div class="flex items-center mb-2 relative">
                <div :class="`w-8 h-8 mr-4 rounded-full ${status.color}`">
                    <span class="text-white text-xs font-semibold flex items-center justify-center h-full" @mouseenter="showTooltip = true" @mouseleave="showTooltip = false">
                        {{ status.percentage.toFixed(0) }}%
                    </span>
                    <span
                        class="absolute top-10 bg-black text-white text-xs rounded px-2 py-1 opacity-0 pointer-events-none transition-opacity duration-200"
                        :class="{ 'opacity-100 pointer-events-auto': showTooltip }"
                    >
                        {{ $t('sensors.list.accuracy_of', { sensor: sensor.name }) }}
                    </span>
                </div>
            <div class="font-semibold text-lg">{{ sensor.name }}</div>
            </div>
            <div class="text-sm text-gray-600 mb-1">
                <div>
                    <span class="font-bold">{{ $t('sensors.list.id') }}: </span>
                    <span>{{ " " + sensor.id }}</span>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-2 mt-2">

                    <div>
                        <span class="font-bold">{{ $t('sensors.list.average') }}: </span>
                        <span>{{ " " + sensor.average.toFixed(2) }}</span> <span class="text-xs">dispmm</span>
                    </div>
                    <div>
                        <span class="font-bold">{{ $t('sensors.list.min') }}:</span>
                        <span>{{ " " + sensor.min.toFixed(2) }}</span> <span class="text-xs">dispmm</span>
                    </div>
                    <div>
                        <span class="font-bold">{{ $t('sensors.list.max') }}:</span>
                        <span>{{ " " + sensor.max.toFixed(2) }}</span> <span class="text-xs">dispmm</span>
                    </div>
                    <div>
                        <span class="font-bold">{{ $t('sensors.list.status') }}:</span>
                        <span>{{ " " + sensor.status }}</span>
                    </div>
                 </div> 
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
import { computed, defineProps, ref } from 'vue';
import Sensor from '../models/sensor';

const props = defineProps<{
    sensor: Sensor;
}>();

const showTooltip = ref(false);
const status = computed<{
    color: string;
    percentage: number;
}>(() => {
    // calculate the amount of measurements going over the threshold
    const overThreshold = props.sensor.measurements.filter(m => m.value > props.sensor.threshold).length;
    const totalMeasurements = props.sensor.measurements.length;
    const percentage = totalMeasurements > 0 ? ((100 - overThreshold) / totalMeasurements) * 100 : 0;
    let color = 'bg-gray-200'; // default color
    if (percentage > 90) {
        color = 'bg-green-500';
    } else if (percentage > 80) {
        color = 'bg-yellow-500';
    } else if (percentage > 25) {
        color = 'bg-red-500';
    }
    return {
        color,
        percentage
    };
})

</script>
<template>
    <div>
        <h1 class="text-2xl font-bold mb-4">{{ $t('sensors.list.title') }}</h1>
        <div class="grid grid-cols-5 gap-4 mb-4">
            <input type="text" v-model="query" placeholder="Search sensors..." class="border border-gray-300 rounded-lg p-2 mb-4 w-full h-12 col-span-4" />
            <button type="button" @click="search" class="bg-blue-500 text-white rounded-lg h-12 col-span-1" value="Search">Search</button>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div v-for="sensor in sensors" :key="sensor.id" class="bg-white shadow p-4 border border-gray-200 rounded-lg grid grid-cols-1 gap-4 
            hover:shadow-lg transition-shadow duration-200 cursor-pointer" @click="() => $router.push({ name: 'sensorDetails', params: { id: sensor.id } })">
                <section>
                    <div class="flex items-center mb-2">
                        <div :class="`w-8 h-8 mr-4 rounded-full ${getStatusColor(sensor)}`"></div>
                    <div class="font-semibold text-lg">{{ sensor.name }}</div>
                    </div>
                    <div class="text-sm text-gray-600 mb-1">
                        <span class="font-medium">{{ $t('sensors.list.id') }}:</span> {{ sensor.id }}
                    </div>
                    <div class="text-sm text-gray-600">
                        <span class="font-medium">{{ $t('sensors.list.location') }}:</span> {{ sensor.location }}
                    </div>
                </section>
                <!-- <section class="flex justify-between items-center">
                    <div ref="graph"></div>
                </section> -->
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import * as am5 from '@amcharts/amcharts5';
import * as am5xy from '@amcharts/amcharts5/xy';
import am5themes_Animated from '@amcharts/amcharts5/themes/Animated';
import { onMounted, ref } from 'vue';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';

const query = ref('');
const graph = ref<HTMLElement | null>(null);
const sensors = ref<Sensor[]>([]);

onMounted(() => {
    sensorsApi.getSensors().then(response => {
        sensors.value = response;
    }).catch(error => {
        console.error('Error fetching sensors:', error);
    });
})

function getStatusColor(sensor: Sensor): string {
    return [
                'bg-green-500 text-gray-800',
                'bg-red-500 text-gray-900',
            ][Math.floor(Math.random() * 2)]
    switch (sensor.status) {
        case 'active':
            return 'bg-green-100 text-green-800';
        case 'inactive':
            return 'bg-yellow-100 text-yellow-800';
        case 'error':
            return 'bg-red-100 text-red-800';
        default:
            return [
                'bg-gray-100 text-gray-800',
                'bg-gray-500 text-gray-900',
                'bg-gray-900 text-gray-900'
            ][Math.floor(Math.random() * 3)];
    }
}

function search() {
    // Implement search logic here
    // For now, just log the query
    console.log('Searching for:', query.value);
}

// function mountGraph() {
//     let root = am5.Root.new(graph.value!);
//     root.setThemes([am5themes_Animated.new(root)]);
//     let chart = root.container.children.push(
//         am5xy.XYChart.new(root, {
//             panY: false,
//             layout: root.verticalLayout
//         })
//     );

//     let data = [{
//         category: "Research",
//         value1: 1000,
//         value2: 588
//       },
//  {
//         category: "Marketing",
//         value1: 1200,
//         value2: 1800
//       }, {
//         category: "Sales",
//         value1: 850,
//         value2: 1230
//       }
//     ];

//     // Create Y-axis
//     let yAxis = chart.yAxes.push(
//       am5xy.ValueAxis.new(root, {
//         renderer: am5xy.AxisRendererY.new(root, {})
//       })
//     );

//     // Create X-Axis
//     let xAxis = chart.xAxes.push(
//       am5xy.CategoryAxis.new(root, {
//         renderer: am5xy.AxisRendererX.new(root, {}),
//         categoryField: "category"
//       })
//     );
//     xAxis.data.setAll(data);

//     // Create series
//     let series1 = chart.series.push(
//       am5xy.ColumnSeries.new(root, {
//         name: "Series",
//         xAxis: xAxis,
//         yAxis: yAxis,
//         valueYField: "value1",
//         categoryXField: "category"
//       })
//     );
//     series1.data.setAll(data);

//     let series2 = chart.series.push(
//       am5xy.ColumnSeries.new(root, {
//         name: "Series",
//         xAxis: xAxis,
//         yAxis: yAxis,
//         valueYField: "value2",
//         categoryXField: "category"
//       })
//     );
//     series2.data.setAll(data);

//     // Add legend
//     let legend = chart.children.push(am5.Legend.new(root, {}));
//     legend.data.setAll(chart.series.values);

//     // Add cursor
//     chart.set("cursor", am5xy.XYCursor.new(root, {}));
// }

onMounted(() => {
});
</script>
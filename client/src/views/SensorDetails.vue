<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import { VueSpinnerBall } from 'vue3-spinners';
import * as am5 from "@amcharts/amcharts5";
import * as am5xy from "@amcharts/amcharts5/xy";
import measurementsApi from '../api/measurements';
import type { Measurement } from '../models/measurement';

let sensor = ref<Sensor | null>(null);
let measurements = ref<Measurement[]>([]);
let startDate = ref<string | null>(null);
let endDate = ref<string | null>(null);
let chartDiv = ref<HTMLElement | null>(null);
const route = useRoute();

function formatDateToInput(date: Date): string {
    return date.toISOString().split('T')[0];
}

onMounted(() => {
    const id = route.params.id;
    sensorsApi.getSensorById(id as string).then(response => {
        sensor.value = response;
        if (sensor.value) {
            measurementsApi.getMeasurements(sensor.value.id).then(m => {
                measurements.value = m;
                sensor.value!.measurements = measurements.value;
                startDate.value = measurements.value.length > 0 ? formatDateToInput(new Date(measurements.value[0].timestamp)) : null;
                endDate.value = measurements.value.length > 0 ? formatDateToInput(new Date(measurements.value[measurements.value.length - 1].timestamp)) : null;
                buildChart(measurements.value);
            }).catch(error => {
                console.error('Error fetching measurements:', error);
            });
        }
    }).catch(error => {
        console.error('Error fetching sensor details:', error);
    });
})

function buildChart(measurements: Measurement[]) {
    if (!chartDiv.value) {
        console.error('Chart div is not available');
        return;
    }

    let root = am5.Root.new(chartDiv.value);
    let chart = root.container.children.push(
        am5xy.XYChart.new(root, {
            panX: true,
            panY: true,
            wheelX: "panX",
            wheelY: "zoomX",
        })
    )
    let yAxis = chart.yAxes.push(
        am5xy.ValueAxis.new(root, {
            min: 0,
            max: 100,
            renderer: am5xy.AxisRendererY.new(root, {
                
            })
        })
    )
    let xAxis = chart.xAxes.push(
        am5xy.DateAxis.new(root, {
            baseInterval: { timeUnit: "day", count: 1 },
            renderer: am5xy.AxisRendererX.new(root, {})
        })
    )

    let series = chart.series.push(
        am5xy.LineSeries.new(root, {
            name: sensor.value?.name || 'Sensor Data',
            xAxis: xAxis,
            yAxis: yAxis,
            valueYField: "value",
            valueXField: "timestamp",
            tooltip: am5.Tooltip.new(root, {
                labelText: "{valueY}"
            })
        })
    );

    series.data.setAll(measurements.map(d => ({
        value: d.value,
        timestamp: new Date(d.timestamp).getTime()
    })));
}

</script>

<template>
    <div class="p-4">
        <h1 class="text-2xl font-bold mb-4">
            <span v-if="sensor">{{ sensor.name }}</span>
            <span v-else>Loading Sensor Details...</span>
        </h1>
        <div v-if="sensor">
            <p><strong>ID:</strong> {{ sensor.id }}</p>
            <p><strong>Name:</strong> {{ sensor.name }}</p>
            <p><strong>Type:</strong> {{ sensor.type }}</p>
            <p><strong>Status:</strong> {{ sensor.status }}</p>
            <p><strong>Number of measurements:</strong> {{ sensor.measurements.length }}</p>
        </div>

        <div v-else class="flex justify-center items-center h-full">
            <VueSpinnerBall size="50" color="red" />
        </div>

        <div class="h-90 bg-white shadow-md rounded-lg mt-4">
            <div ref="chartDiv" class="h-full w-full">
            </div>
        </div>

        <div class="flex justify-center items-center h-full mt-4 shadow-md rounded-lg p-4">
            <label>From:
                <input
                    type="date"
                    v-model="startDate"
                    class="ml-2 p-1 border rounded"
                    :min="formatDateToInput(new Date(Date.now() - 30 * 24 * 60 * 60 * 1000))"
                    :max="endDate ? endDate : formatDateToInput(new Date())"
                />
            </label>
            <label class="ml-4">To:
                <input type="date"
                :min="startDate ? startDate : formatDateToInput(new Date(Date.now() - 30 * 24 * 60 * 60 * 1000))"
                :max="formatDateToInput(new Date())"
                v-model="endDate" class="ml-2 p-1 border rounded" />
            </label>
        </div>
    </div>
</template>
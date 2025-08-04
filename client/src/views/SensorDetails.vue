<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Sensor from '../models/sensor';
import sensorsApi from '../api/sensors';
import { VueSpinnerBall } from 'vue3-spinners';
import * as am5 from "@amcharts/amcharts5";
import * as am5xy from "@amcharts/amcharts5/xy";
import measurementsApi from '../api/measurements';
import type { Measurement } from '../models/measurement';
import { watch } from 'vue';
import AddMeasurement from '../components/AddMeasurement.vue';

let sensor = ref<Sensor | null>(null);
let coordinates = computed<[number, number] | null>(() => {
    if (sensor.value && sensor.value.location) {
        return [
            Number(sensor.value.location[0]),
            Number(sensor.value.location[1])
        ];
    }
    return null;
});
let measurements = ref<Measurement[]>([]);
let startDate = ref<string | null>(null);
let endDate = ref<string | null>(null);
let graphStartDate = computed(() => {
    if (startDate.value) {
        const date = new Date(startDate.value);
        date.setDate(date.getDate() + 1);
        return date;
    }
    return new Date(Date.now() - 30 * 24 * 60 * 60 * 1000); // Default to 30 days ago
});
let graphEndDate = computed(() => {
    if (endDate.value) {
        return new Date(endDate.value);
    }
    return new Date(); // Default to today
});
let chartDiv = ref<HTMLElement | null>(null);
let accuracy = computed(() => {
    if (sensor.value && sensor.value.measurements.length > 0) {
        // check the percentage of measurements that are within the threshold
        const overThreshold = sensor.value.measurements.filter(m => m.value > sensor.value?.threshold!).length;
        const totalMeasurements = sensor.value.measurements.length;
        if (totalMeasurements > 0) {
            return ((totalMeasurements - overThreshold) / totalMeasurements) * 100;
        }
    }
    return null;
});
const route = useRoute();
const router = useRouter();

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

let chartRoot: am5.Root | null = null;
let chart: am5xy.XYChart | null = null;
let xAxis: am5xy.DateAxis<am5xy.AxisRenderer> | null = null;
let yAxis: am5xy.ValueAxis<am5xy.AxisRenderer> | null = null;
let series: am5xy.LineSeries | null = null;

function buildChart(measurements: Measurement[]) {
    if (!chartDiv.value) {
        console.error('Chart div is not available');
        return;
    }

    // Dispose previous chart if exists
    if (chartRoot) {
        chartRoot.dispose();
    }

    chartRoot = am5.Root.new(chartDiv.value);
    chart = chartRoot.container.children.push(
        am5xy.XYChart.new(chartRoot, {
            panX: true,
            panY: true,
            wheelX: "panX",
            wheelY: "zoomX",
        })
    )
    yAxis = chart.yAxes.push(
        am5xy.ValueAxis.new(chartRoot, {
            min: sensor.value?.measurements.length ? Math.min(...measurements.map(m => m.value)) : 0,
            max: sensor.value?.measurements.length ? Math.max(...measurements.map(m => m.value)) : 100,
            renderer: am5xy.AxisRendererY.new(chartRoot, {})
        })
    )
    xAxis = chart.xAxes.push(
        am5xy.DateAxis.new(chartRoot, {
            min: new Date(graphStartDate.value).getTime(),
            max: new Date(graphEndDate.value).getTime(),
            baseInterval: { timeUnit: "second", count: 1 },
            renderer: am5xy.AxisRendererX.new(chartRoot, {})
        })
    )

    series = chart.series.push(
        am5xy.LineSeries.new(chartRoot, {
            name: sensor.value?.name || 'Sensor Data',
            xAxis: xAxis,
            yAxis: yAxis,
            valueYField: "value",
            valueXField: "timestamp",
            tooltip: am5.Tooltip.new(chartRoot, {
                labelText: "{valueY}"
            }),
            tooltipHTML: "{valueY}"
        })
    );

    series.data.setAll(measurements.map(d => ({
        value: d.value,
        timestamp: new Date(d.timestamp).getTime()
    })));

    // add an horizontal line at y = 5
    let rangeDataItem = yAxis.makeDataItem({
        value: sensor.value?.threshold || 5, // Default threshold if not set
    });

    // Create a range on series
    let range = series.createAxisRange(rangeDataItem);

    range.series?.template?.setAll({
        stroke: am5.color(0xff0000),
        fill: am5.color(0xff0000)
    });

    rangeDataItem.get("grid")?.setAll({
        stroke: am5.color(0xff0000),
        strokeWidth: 10,
        visible: true
    });
}

// Watch for changes in graphStartDate and graphEndDate to update xAxis min/max
watch([graphStartDate, graphEndDate], ([newStart, newEnd]) => {
    if (xAxis) {
        xAxis.set("min", newStart.getTime());
        xAxis.set("max", newEnd.getTime());
    }
});

function addMeasurement(newMeasurement: Measurement) {
    if (sensor.value) {
        measurementsApi.addMeasurement(sensor.value.id, newMeasurement).then(measurement => {
            measurements.value.push(measurement);
            sensorsApi.getSensorById(sensor.value?.id!).then(response => {
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
            buildChart(measurements.value);
        }).catch(error => {
            console.error('Error adding measurement:', error);
        });
    }
}

function deleteSensor() {
    if (!confirm('Are you sure you want to delete this sensor? This action cannot be undone.')) {
        return;
    }

    if (sensor.value) {
        sensorsApi.deleteSensor(sensor.value.id).then(() => {
            // Redirect to the sensors list or home page after deletion
            router.push('/sensors');
        }).catch(error => {
            console.error('Error deleting sensor:', error);
        });
    }
}

</script>

<template>
    <div class="">
        <h1 class="text-2xl font-bold mb-4">
            <span v-if="sensor">{{ sensor.name }}</span>
            <span v-else>Loading Sensor Details...</span>
        </h1>
        <div v-if="sensor">
            <p><strong>ID:</strong> {{ sensor.id }}</p>
            <p><strong>Name:</strong> {{ sensor.name }}</p>
            <p><strong>Status:</strong> {{ sensor.status() }}</p>
            <p><strong>Number of measurements:</strong> {{ sensor.measurements.length }}</p>
            <p><strong>Threshold:</strong> {{ sensor.threshold.toPrecision(5) }}</p>
            <p><strong>Accuracy:</strong> {{ accuracy ? accuracy.toFixed(2) : 'N/A' }}%</p>
        </div>

        <div v-else class="flex justify-center items-center h-full">
            <VueSpinnerBall size="50" color="red" />
        </div>

        <div v-if="coordinates" class="mt-4">
            <ol-map style="min-width: 400px; height: 400px;">
                <ol-view :center="coordinates" :zoom="5" projection="EPSG:4326" />
                <ol-tile-layer>
                    <ol-source-osm />
                </ol-tile-layer>
                <ol-vector-layer>
                    <ol-source-vector>
                        <ol-feature>
                            <ol-geom-point :coordinates="coordinates"></ol-geom-point>
                            <ol-style>
                                <ol-style-circle :radius="1">
                                    <ol-style-fill color="red"></ol-style-fill>
                                    <ol-style-stroke color="red" width="10"></ol-style-stroke>
                                </ol-style-circle>
                            </ol-style>
                        </ol-feature>
                    </ol-source-vector>
                </ol-vector-layer>
            </ol-map>
        </div>

        <div class="h-90 bg-white shadow-md rounded-lg mt-4">
            <div ref="chartDiv" class="h-full w-full">
            </div>
        </div>

        <div class="flex justify-center items-center h-full mt-4 shadow-md rounded-lg p-4">
            <label>From:
                <input type="date" v-model="startDate" class="ml-2 p-1 border rounded"
                    :min="formatDateToInput(new Date(Date.now() - 30 * 24 * 60 * 60 * 1000))"
                    :max="endDate ? endDate : formatDateToInput(new Date())" />
            </label>
            <label class="ml-4">To:
                <input type="date"
                    :min="startDate ? startDate : formatDateToInput(new Date(Date.now() - 30 * 24 * 60 * 60 * 1000))"
                    :max="formatDateToInput(new Date(Date.now() + 24 * 60 * 60 * 1000))" v-model="endDate" class="ml-2 p-1 border rounded" />
            </label>
        </div>

        <AddMeasurement v-if="sensor" :sensorId="sensor.id" @addMeasurement="addMeasurement" />

        <div class="flex justify-center items-center h-full mt-4">
            <button v-if="sensor" @click="deleteSensor" class="mt-4 px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-opacity-50">
                {{ $t('sensors.create.delete_button') }}
            </button>
        </div>
    </div>
</template>
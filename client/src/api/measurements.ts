import { Measurement } from "../models/measurement";
import ApiService from "./server";

class MeasurementsApiService extends ApiService<Measurement> {
    constructor() {
        super({ path: "measurements", modelClass: Measurement });
    }

    async getMeasurements(sensorId: string): Promise<Measurement[]> {
        const measurements = await this.fetchData({ query: { deviceId: sensorId } });
        return measurements; // Assuming the API returns an array of measurement objects
    }

    async addMeasurement(sensorId: string, measurement: Measurement): Promise<Measurement> {
        const response = await this.fetchData({
            method: "POST",
            body: [measurement.value],
            endpoint: sensorId
        });
        return response[0]; // Assuming the API returns the created measurement object
    }
}

const measurementsApi = new MeasurementsApiService();
export default measurementsApi;

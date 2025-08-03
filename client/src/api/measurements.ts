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
}

const measurementsApi = new MeasurementsApiService();
export default measurementsApi;

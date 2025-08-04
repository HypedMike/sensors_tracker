import Sensor from "../models/sensor";
import ApiService from "./server";

class SensorsApiService extends ApiService<Sensor> {
    constructor() {
        super({ path: "devices", modelClass: Sensor });
    }

    async getSensors(full: boolean = false): Promise<Sensor[]> {
        return this.fetchData({
            query: full ? { m: 'true' } : undefined,
            method: 'GET'
        });
    }

    async getSensorById(id: string): Promise<Sensor> {
        const sensors = await this.fetchData({ endpoint: id });
        return sensors[0]; // Assuming the API returns a single sensor object
    }

    async createSensor(sensor: Sensor): Promise<any> {
        return this.fetchData({
            method: 'POST',
            body: {
                ...sensor,
                location: [
                    Number(sensor.location[0]),
                    Number(sensor.location[1])
                ]
            }
        });
    }

    async deleteSensor(id: string): Promise<void> {
        await this.fetchData({
            method: 'DELETE',
            endpoint: id
        });
    }
}

const sensorsApi = new SensorsApiService();
export default sensorsApi;

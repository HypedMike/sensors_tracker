import Sensor from "../models/sensor";
import ApiService from "./server";

class SensorsApiService extends ApiService<Sensor> {
    constructor() {
        super({ path: "devices", modelClass: Sensor });
    }
    
    async getSensors(): Promise<Sensor[]> {
        return this.fetchData();
    }

    async getSensorById(id: string): Promise<Sensor> {
        const sensors = await this.fetchData({ endpoint: id });
        return sensors[0]; // Assuming the API returns a single sensor object
    }
}

const sensorsApi = new SensorsApiService();
export default sensorsApi;

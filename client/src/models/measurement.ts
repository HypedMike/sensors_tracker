import type { Model } from "./common";

export class Measurement implements Model{
    id: string;
    sensorId: string;
    value: number;
    timestamp: Date;

    constructor(props: Partial<Measurement>) {
        this.id = props.id || '';
        this.sensorId = props.sensorId || '';
        this.value = props.value || 0;
        this.timestamp = props.timestamp ? new Date(props.timestamp) : new Date();
    }

    fromJSON(json: any): this {
        this.id = json.id;
        this.sensorId = json.sensorId;
        this.value = json.value;
        this.timestamp = new Date(json.timestamp);
        return this;
    }
}
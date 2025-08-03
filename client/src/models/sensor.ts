import type { Model } from "./common";
import { Measurement } from "./measurement";

export default class Sensor implements Model {
  id: string;
  name: string;
  location: [number, number];
  threshold: number;
  status: string;
  createdAt: Date;
  updatedAt: Date;

  measurements: Measurement[];
  min: number;
  max: number;
  average: number;

  constructor(props: Partial<Sensor>) {
    this.id = props.id || '';
    this.name = props.name || '';
    this.location = props.location || [0, 0];
    this.status = props.status || 'active';
    this.createdAt = props.createdAt ? new Date(props.createdAt) : new Date();
    this.updatedAt = props.updatedAt ? new Date(props.updatedAt) : new Date();
    this.measurements = props.measurements ? props.measurements.map(m => new Measurement(m)) : [];
    this.threshold = props.threshold || 0;
    this.min = props.min || 0;
    this.max = props.max || 0;
    this.average = props.average || 0;
  }

  fromJSON(json: any): this {
    this.id = json.id;
    this.name = json.name;
    this.location = json.location;
    this.status = json.status;
    this.createdAt = new Date(json.createdAt);
    this.updatedAt = new Date(json.updatedAt);
    this.measurements = json.measurements ? json.measurements.map((m: any) => new Measurement(m)) : [];
    this.threshold = json.threshold || 0;
    this.min = json.min || 0;
    this.max = json.max || 0;
    this.average = json.average || 0;
    return this;
  }
}
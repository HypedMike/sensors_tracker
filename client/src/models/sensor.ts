import type { Model } from "./common";
import { Measurement } from "./measurement";

export default class Sensor implements Model {
  id: string;
  name: string;
  type: string;
  location: string;
  status: string;
  createdAt: Date;
  updatedAt: Date;

  measurements: Measurement[];

  constructor(props: Partial<Sensor>) {
    this.id = props.id || '';
    this.name = props.name || '';
    this.type = props.type || '';
    this.location = props.location || '';
    this.status = props.status || 'active';
    this.createdAt = props.createdAt ? new Date(props.createdAt) : new Date();
    this.updatedAt = props.updatedAt ? new Date(props.updatedAt) : new Date();
    this.measurements = props.measurements ? props.measurements.map(m => new Measurement(m)) : [];
  }

  fromJSON(json: any): this {
    this.id = json.id;
    this.name = json.name;
    this.type = json.type;
    this.location = json.location;
    this.status = json.status;
    this.createdAt = new Date(json.createdAt);
    this.updatedAt = new Date(json.updatedAt);
    return this;
  }
}
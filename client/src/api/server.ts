import type { Model } from "../models/common";

const baseUrl = import.meta.env.VITE_API_BASE_URL;

abstract class ApiService<T extends Model> {
  protected baseUrl: string;
  protected path: string;
  protected modelClass: { new(props: any): T };

  constructor({
    path = "",
    modelClass
  }: {
    path?: string;
    modelClass: { new(props: any): T };
  }) {
    this.baseUrl = baseUrl;
    this.path = path;
    this.modelClass = modelClass;
  }

  protected async fetchData(options?: { endpoint?: string; method?: string; body?: any, query?: Record<string, any> }): Promise<T[]> {
    try {
      const response = await fetch(`${this.baseUrl}/${this.path}${options?.endpoint ? '/' + options.endpoint : ''}?${new URLSearchParams(options?.query).toString()}`, {
        method: options?.method || "GET",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(options?.body),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();

      // check if data is an array
      if (Array.isArray(data)) {
        return data.map((item: any) => new this.modelClass(item));
      }

      // if data is not an array, treat it as a single object
      return [new this.modelClass(data)];

    } catch (error) {
      console.error("Fetch error:", error);
      throw error;
    }
  }
}

export default ApiService;

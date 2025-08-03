import type { Model } from "./common";

export class User implements Model {
    id: string;
    username: string;
    token: string;

    constructor(props: Partial<User>) {
        this.id = props.id || "";
        this.username = props.username || "";
        this.token = props.token || "";
    }

    fromJSON(json: any): this {
        return new User({
            id: json.id,
            username: json.username,
            token: json.token
        }) as this;
    }
}
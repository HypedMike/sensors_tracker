import { User } from "../models/user";
import ApiService from "./server";

class AuthApiService extends ApiService<User> {
    constructor() {
        super({ path: "auth", modelClass: User });
    }

    async login(username: string, password: string) {
        const res = await this.fetchData({
            endpoint: "login",
            method: "POST",
            body: {
                username,
                password
            }
        })  
        
        return res[0];
    }

    async getUser() {
        const res = await this.fetchData({
            endpoint: "user",
            method: "GET"
        });

        return res[0];
    }

}

const authService = new AuthApiService();
export default authService;

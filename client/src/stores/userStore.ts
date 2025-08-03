import { defineStore } from "pinia";
import { User } from "../models/user";

export const useUserStore = defineStore('userStore', {
    state: () => ({
        user: new User({})
    }),
    getters: {
        isLoggedIn: (state) => {
            return !!state.user.token;
        }
    },
    actions: {
        setUser(user: User) {
            this.user = user;
        },
        clearUser() {
            this.user = new User({});
        },
        updateUser(user: Partial<User>) {
            this.user = { ...this.user, ...user };
        }
    }
});
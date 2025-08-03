<template>
    <div class="flex items-center justify-center h-screen">
        <section class="grid grid-cols-1 gap-4 w-96 p-6 bg-white shadow-md rounded-lg">
            <h2 class="text-2xl font-bold mb-4">{{ $t('login.title') }}</h2>
            <input type="text" v-model="user.username" placeholder="Username" />
            <input type="password" v-model="password" placeholder="Password" />
            <button type="submit" @click="onSubmit">Login</button>
            <div class="text-red-500" v-if="error && error.length > 0">
                {{ error }}
            </div>
        </section>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import authService from '../api/auth';
import { User } from '../models/user';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/userStore';

let user = ref<User>(new User({}))
let password = ref<string>("")
let error = ref<string | null>("");

const router = useRouter();
const userStore = useUserStore();

function onSubmit() {
    authService.login(user.value.username, password.value).then((e) => {
       document.cookie = `token= ${e.token}; path=/; max-age=31536000`;
        router.replace("/");
        // update store
        userStore.setUser(new User({
            username: user.value.username,
            token: e.token
        }));
    }).catch((e) => {
        console.log(e);
        error.value = e.toString();
    })
}
</script>
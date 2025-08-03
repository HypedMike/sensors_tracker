<script setup lang="ts">
import { RouterView } from 'vue-router';
import Navbar from './components/Navbar.vue';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from './stores/userStore';
import authService from './api/auth';
import { User } from './models/user';

let mobileScreen = ref<boolean>(window.innerWidth < 768);
let showNavbar = ref<boolean>(window.innerWidth >= 768);


const router = useRouter();
const userStore = useUserStore();

router.afterEach(() => {
    // Close the navbar on route change if on mobile
    if (mobileScreen.value) {
        showNavbar.value = false;
    }
});

onMounted(() => {
    // Check if the user is logged in
    if (!userStore.isLoggedIn) {
      // check if there is a JWT token in cookies
      const cookies = document.cookie.split(';').find(cookie => cookie.trim().startsWith('token='));
      if (cookies) {
        authService.getUser().then(user => {
          console.log('User fetched from API:', user);
          userStore.setUser(new User({
            username: user.username,
            token: cookies.split('=')[1] // Extract the token from the cookie
          }));
          router.push('/');
        }).catch(() => {
          // If the token is invalid, redirect to login and clear the store
          userStore.clearUser();
          document.cookie = 'token=; path=/; max-age=0'; // Clear the token cookie
          router.push('/login');
        });
      } else {
        // If no token is found, redirect to login
        router.push('/login');
      }
    }
});

</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-[300px_1fr] h-screen">
    <Navbar v-if="showNavbar" @close="showNavbar = false" />
    <button v-if="mobileScreen" class="fixed top-4 left-4 z-50" @click="showNavbar = !showNavbar">
      <img src="/icons/menu.png" alt="Menu" class="h-6 w-6" />
    </button>
    <div :class="`p-4 bg-gray-100 overflow-y-auto h-full ${mobileScreen ? 'pt-10' : ''}`">
      <RouterView />
    </div>
  </div>
</template>

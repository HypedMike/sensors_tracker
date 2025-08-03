<script setup lang="ts">
import { RouterView } from 'vue-router';
import Navbar from './components/Navbar.vue';

</script>

<template>
  <div class="grid grid-cols-[300px_1fr] h-screen">
    <Navbar />
    <div class="p-4 bg-gray-100 overflow-y-auto h-full">
      <RouterView />
    </div>
  </div>
</template>

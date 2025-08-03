<template>

</template>

<script setup lang="ts">
import { onMounted } from 'vue';

onMounted(() => {
    document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
})
</script>
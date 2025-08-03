<template>
    <nav class="flex flex-col w-[300px] h-full border-r border-gray-200">
        <section class="h-[200px] flex items-center justify-center ">
            <RouterLink class="p-4 flex items-center rounded-lg" to="/">
                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQLGrVXxH4Nhcrk9cO5rY10pXMe7310EAwCkQ&s" alt="Logo" class="h-8" />
            </RouterLink>
        </section>
        <section class="flex flex-col items-center gap-4 p-4">
            {{ $t('navbar.choose_language') }}
            <select v-model="$i18n.locale">
                <option v-for="locale in $i18n.availableLocales" :key="`locale-${locale}`" :value="locale">{{ locale }}</option>
            </select>
        </section>
        <section class="flex flex-col gap-4 items-center h-full">
            <RouterLink v-for="route in routes" :key="route.name" 
            style="background-color: var(--secondary);"
                :to="route.path" class="w-[70%] hover:text-gray-300 underline text-center p-1 rounded-lg text-white drop-shadow-lg">
                {{ route.name }}
            </RouterLink>
        </section>
    </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { ref } from 'vue';

const { t } = useI18n();

import { watchEffect } from 'vue';

const routes = ref<{
    path: string;
    name: string;
    sort: number;
}[]>([]);

watchEffect(() => {
    routes.value = [
        {
            path: '/sensors',
            name: t('sensors.title'),
            sort: 1
        },
        {
            path: '/logout',
            name: t('logout.title'),
            sort: 3
        },
        {
            path: '/sensors/new',
            name: t('sensors.create.title'),
            sort: 2
        }
    ].sort((a, b) => a.sort - b.sort);
});

</script>
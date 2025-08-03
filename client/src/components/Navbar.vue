<template>
    <nav
        class="flex flex-col w-[300px] h-full w-full border-r border-gray-200
                     fixed top-0 left-0 z-50 bg-white
                     md:static md:w-[300px] md:h-full
                     min-h-screen
                     transition-all
                     sm:w-full sm:h-screen sm:border-r-0 sm:border-b sm:fixed sm:top-0 sm:left-0"
        style="max-width:100vw;"
    >
        <section class="h-[200px] flex items-center justify-center">
            <RouterLink class="p-4 flex items-center rounded-lg" to="/">
                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQLGrVXxH4Nhcrk9cO5rY10pXMe7310EAwCkQ&s" alt="Logo" class="h-8" />
            </RouterLink>
        </section>
        <section class="flex flex-col items-center gap-4 p-4">
            {{ $t('navbar.choose_language') }}
            <select v-model="$i18n.locale" @change="setLanguagePreference($i18n.locale)">
                <option v-for="locale in $i18n.availableLocales" :key="`locale-${locale}`" :value="locale">{{ locale }}</option>
            </select>
        </section>
        <section class="flex flex-col gap-4 items-center h-full" v-if="userStore.isLoggedIn">
            <RouterLink v-for="route in routes" :key="route.name"
                style="background-color: var(--secondary);"
                :to="route.path" class="w-[70%] hover:text-gray-300 underline text-center p-1 rounded-lg text-white drop-shadow-lg">
                {{ route.name }}
            </RouterLink>
        </section>
        <section class="w-full flex justify-center items-center p-4 md:hidden">
            <button>
            <img src="/icons/close.png" alt="Close" class="h-6 w-6" @click="$emit('close')" />
            </button>
        </section>
    </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { onMounted, ref } from 'vue';

const { t, locale } = useI18n();

import { watchEffect } from 'vue';
import { useUserStore } from '../stores/userStore';

const emit = defineEmits(['close']);

const routes = ref<{
    path: string;
    name: string;
    sort: number;
}[]>([]);

const userStore = useUserStore();

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

function checkLanguagePreference() {
    // get cookies
    const cookies = document.cookie.split(';').reduce((acc, cookie) => {
        const [key, value] = cookie.trim().split('=');
        acc[key] = value;
        return acc;
    }, {} as Record<string, string>);

    // check if language is set
    if (cookies['preferredLanguage']) {
        // set the language
        locale.value = cookies['preferredLanguage'];
    } else {
        // set the default language
        locale.value = 'en';
    }
}

function setLanguagePreference(lang: string) {
    // set the cookie
    document.cookie = `preferredLanguage=${lang}; path=/; max-age=31536000`;
}

onMounted(() => {
    checkLanguagePreference();
})

</script>
import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import HomeView from './views/HomeView.vue'
import { createRouter, createWebHistory } from 'vue-router'
import SensorsListView from './views/SensorsListView.vue'
import CreateSensor from './views/CreateSensor.vue'
import { createI18n } from 'vue-i18n'
import * as IT from './lang/it'
import * as EN from './lang/en'
import SensorDetails from './views/SensorDetails.vue'
import Logout from './views/Logout.vue'
import "vue3-openlayers/styles.css";
import OpenLayersMap from "vue3-openlayers";
import LoginView from './views/LoginView.vue'
import { createPinia } from 'pinia'



const routes = [
    {
        path: '/',
        component: HomeView,
    },
    {
        path: '/sensors',
        component: SensorsListView,
    },
    {
        path: '/sensors/new',
        component: CreateSensor
    },
    {
        path: '/logout',
        component: Logout,
    },
    {
        path: '/sensors/:id',
        name: 'sensorDetails',
        component: SensorDetails
    },
    {
        path: '/login',
        name: 'login',
        component: LoginView
    }
]

const i18n = createI18n({
    legacy: false,
  locale: 'en',
  messages: {
    it: IT.default,
    en: EN.default
  }
});

const router = createRouter({
    history: createWebHistory(),
    routes
});

const pinia = createPinia();

createApp(App).use(router).use(i18n).use(OpenLayersMap).use(pinia).mount('#app')

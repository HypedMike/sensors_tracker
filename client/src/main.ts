import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import HomeView from './views/HomeView.vue'
import { createRouter, createWebHistory } from 'vue-router'
import SensorsListView from './views/SensorsListView.vue'
import Logout from './components/Logout.vue'
import CreateSensor from './views/CreateSensor.vue'
import { createI18n } from 'vue-i18n'
import * as IT from './lang/it'
import * as EN from './lang/en'
import SensorDetails from './views/SensorDetails.vue'


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
    }
]

const i18n = createI18n({
    legacy: false,
  locale: 'en',
  messages: {
    it: IT.default,
    en: EN.default
  }
})

const router = createRouter({
    history: createWebHistory(),
    routes
})

createApp(App).use(router).use(i18n).mount('#app')

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import './assets/main.css'
import router from './router'
import PrimeVue from 'primevue/config'
import 'primeicons/primeicons.css'
import Aura from '@primevue/themes/aura'
const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
})
app.mount('#app')
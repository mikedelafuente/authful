import { createApp } from 'vue'
import VueCookies from 'vue3-cookies'
import App from './App.vue'
import router from './router'

 createApp(App).use(router).use(VueCookies).mount('#app')
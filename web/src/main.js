import { createApp } from 'vue'
import VueCookies from 'vue3-cookies'
import App from './App.vue'
import router from './router'
// import "bootstrap/dist/css/bootstrap.min.css"
import 'bootstrap'

 createApp(App).use(router).use(VueCookies).mount('#app')
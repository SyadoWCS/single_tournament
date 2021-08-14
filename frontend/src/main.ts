import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

axios.defaults.baseURL = process.env.NODE_ENV === 'production' ? 'https://backend.syadowcs-tournament.com/api/' : 'http://localhost/api/'
axios.defaults.withCredentials = true

createApp(App).use(store).use(router).mount('#app')

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import './assets/css/global.css'
import config from '../config/config'
import axios from 'axios'
import 'mavon-editor/dist/css/index.css'
import mavonEditor from 'mavon-editor'
import globalVal from './global/globalVal.vue'

Vue.config.productionTip = false
Vue.prototype.GLOBAL = globalVal
Vue.use(ElementUI);
Vue.use(mavonEditor);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

axios.defaults.baseURL = config.host

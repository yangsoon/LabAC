import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'

Vue.config.productionTip = false;

Vue.use(ElementUI);

import store from "./store/main"
import router from "./router/main"
import "./assets/iconfont"

new Vue({
  render: h => h(App),
  router,
  store: store,
}).$mount('#app');

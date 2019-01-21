import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import App from './App.vue';
import VueRouter from 'vue-router';
import store from './store';
import router from './routes';

Vue.use(VueRouter);
Vue.use(Vuetify);

new Vue({
  router,
  store,
  el: '#app',
  render: h => h(App)
})

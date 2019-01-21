import VueRouter from 'vue-router';
import Account from './components/account.vue';
import Campaigns from './components/campaigns.vue';

const routes = [
  { path: '/', redirect: '/campaigns' },
  { path: '/account', component: Account },
  { path: '/campaigns', component: Campaigns }
]

const router = new VueRouter({
  routes,
  mode: 'history',
})

export default router;

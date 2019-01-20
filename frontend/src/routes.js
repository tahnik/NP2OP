import VueRouter from 'vue-router';
import SignIn from './components/sign_in.vue';
import Register from './components/register.vue';

const routes = [
  { path: '/sign-in', component: SignIn },
  { path: '/register', component: Register }
]

const router = new VueRouter({
  routes,
  mode: 'history',
})

export default router;

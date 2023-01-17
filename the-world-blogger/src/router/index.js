import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import UserData from '../views/UserDataView.vue'
import RegisterView from '../views/RegisterView.vue'
import MarkDown from '../views/MarkDown.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: LoginView
  },
  {
    path: '/user/data',
    name: 'userData',
    component: UserData
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/blog/content',
    component: MarkDown
  }
]

//
const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

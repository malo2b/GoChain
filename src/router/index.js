import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import TxIndexView from "@/views/transaction/TxIndexView";
import TxCreateView from "@/views/transaction/TxCreateView";

const routes = [
  /*
  {
    path: '/',
    name: 'home',
    component: HomeView
    },
   */
  {
    path: '/',
    name: 'transaction',
    component: TxIndexView
  },
  {
    path: '/transaction/create',
    name: 'createTransaction',
    component: TxCreateView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

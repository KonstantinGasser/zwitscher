import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/timeline/:id',
    name: 'Timeline',
    component: () => import('../views/Timeline.vue')
  },
  {
    path: '/connections/:id',
    name: 'Connections',
    component: () => import('../views/Connections.vue')
  },
  {
    path: '/tweets/:id',
    name: 'Tweets',
    component: () => import('../views/Tweets.vue')
  },
  {
    path: '/new/:id',
    name: 'NewTweet',
    component: () => import('../views/NewTweet.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

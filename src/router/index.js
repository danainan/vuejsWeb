import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/Login.vue'
import AppView from '../App.vue'
import RegisterView from '../views/Register.vue'



const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '',
      name: 'home',
      component: AppView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    


  ]
})

export default router

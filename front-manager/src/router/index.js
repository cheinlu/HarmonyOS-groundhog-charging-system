import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'hash',
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', component: () => import('@/views/Login') },
    {
      path: '/home',
      component: () => import('@/views/Home'),
      redirect: '/home/dashboard',
      children: [
        { path: 'charge', component: () => import('@/views/charges/Charge') },
        { path: 'pile', component: () => import('@/views/charges/Pile') },
        { path: 'chargeOrder', component: () => import('@/views/orders/ChargeOrder') },
        { path: 'userList', component: () => import('@/views/users/UserList') },
        { path: 'record', component: () => import('@/views/users/Record') },
        { path: 'priceList', component: () => import('@/views/price/priceList') },
        { path: 'dashboard', component: () => import('@/views/dashboard') }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.path === '/home') {
    const token = localStorage.getItem('token')
    if (token) {
      next()
    } else {
      next('/login')
    }
  } else {
    next()
  }
})

export default router

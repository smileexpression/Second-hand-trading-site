import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login/index.vue'
import Layout from '@/views/Layout/index.vue'
import Home from '@/views/Home/index.vue'
import Category from '@/views/Category/index.vue'
import Detail from '@/views/Detail/index.vue'
import Register from '@/views/Register/index.vue'
import Checkout from '@/views/Checkout/index.vue'
import Pay from '@/views/Pay/index.vue'
import PayBack from '@/views/Pay/PayBack.vue'
import Member from '@/views/Member/index.vue'
import UserInfo from '@/views/Member/components/UserInfo.vue'
import UserOrder from '@/views/Member/components/UserOrder.vue'
import SoldOrder from '@/views/Member/components/SoldOrder.vue'

import Release from '@/views/Release/index.vue'
import CartList from '@/views/CartList/index.vue'
import Chat from '@/views/Chat/index.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Layout,
      children: [
        {
          path: '',
          component: Home
        },
        {
          path: 'category/:id',
          component: Category
        },
        {
          path: 'detail/:id',
          component: Detail
        },
        {
          path: 'checkout',
          component: Checkout
          // 还没绑定路由跳转，在P87
        },
        {
          path: 'pay',
          component: Pay
        },
        {
          path: 'paycallback',
          component: PayBack
        },
        {
          path: 'member',
          component: Member,
          children: [
            {
              path: '',
              component: UserInfo
            },
            {
              path: 'order',
              component: UserOrder
            },
            {
              path: 'soldorder',
              component: SoldOrder
            }
          ]
        },
        {
          path: 'release',
          component: Release
        },
        {
          path: 'cartList',
          component: CartList
        },
        {
          path: 'chat',
          component: Chat
        }
      ]
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '/register',
      component: Register
    }
  ],
  // 路由滚动行为定制
  scrollBehavior() {
    return {
      top: 0
    }
  }
})

export default router

import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      component: () => import('@/components/layout/index.vue'),
      children: [
        {
          path: 'example',
          name: '',
          meta: { title: "範例" },
          children: [
            {
              path: 'example1',
              name: '',
              meta: { title: "範例1" },
              component: () => import('@/views/example1/index.vue'),
            }
          ]
        },
        {
          path: 'product',
          name: '',
          meta: { title: "產品" },
          children: [
            {
              path: 'order',
              name: '',
              meta: { title: "訂單" },
              component: () => import('@/views/order/index.vue'),
            }
          ]
        },

      ]

    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/index.vue')
    },

  ]
})


export default router

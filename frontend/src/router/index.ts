import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import Transaction from '../components/Transaction.vue'
import Item from '../components/Item.vue'
import Project from '../components/Project.vue'
import component from '../views/Home.vue'
import TransactionDetail from '../components/TransactionDetail.vue'
import ProjectDetail from '../components/ProjectDetail.vue'
import ItemDetail from '../components/ItemDetail.vue'
import StockIn from '../components/StockIn.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Transaction',
    component: Transaction
  },
  {
    path: '/transactions/:id',
    name: 'Transaction Detail',
    component: TransactionDetail
  },
  {
    path: '/projects',
    name: 'Project',
    component: Project
  },
  {
    path: '/projects/:id',
    name: 'Project Detail',
    component: ProjectDetail
  },
  {
    path: '/items',
    name: 'Item',
    component: Item
  },
  {
    path: '/items/:id',
    name: 'Item Detail',
    component: ItemDetail
  },
  {
    path: '/stockins',
    name: 'Stock In',
    component: StockIn
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

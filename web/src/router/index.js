import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [{
        path: '/',
        name: 'home',
        component: HomeView
    },
    {
        path: '/problemset',
        name: 'problemset',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/ProblemsetView.vue')
    },
    {
        path: '/category',
        name: 'category',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/RanklistView.vue')
    },
    {
        path: '/contest',
        name: 'contest',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/RanklistView.vue')
    },
    {
        path: '/status',
        name: 'status',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
    {
        path: '/ranklist',
        name: 'ranklist',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/RanklistView.vue')
    },
    {
        path: '/about',
        name: 'about',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
    {
        path: '/account',
        name: 'account',
        component: () =>
            import ( /* webpackChunkName: "about" */ '../views/RanklistView.vue')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
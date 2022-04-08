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
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
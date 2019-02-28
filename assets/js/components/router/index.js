import Router from 'vue-router'
import NotFound from '../pages/NotFound.vue'
import home from '../pages/home.vue'
import profile from '../pages/profile.vue'

export default new Router({
    routes: [
        {
            path: '*',
            component: NotFound
        },
        {
            path: '/',
            redirect: '/home'
        },
        {
            path: '/home',
            name: 'Home',
            show_in_nav: true,
            component: home
        },

        {
            path: '/toto/list',
            name: 'Meine Aufgaben',
            show_in_nav: true,
            component: home
        },

        {
            path: '/product/list',
            name: 'Produkte',
            show_in_nav: true,
            component: home
        },

        {
            path: '/timesheet',
            name: 'Zeitplanung',
            show_in_nav: true,
            component: home
        },

        {
            path: '/profile',
            name: 'Profil',
            show_in_nav: false,
            component: profile
        },
    ]
})
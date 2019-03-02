import Vue from 'vue/dist/vue.esm'
import Router from 'vue-router'
import NotFound from '../components/pages/NotFound.vue'
import home from '../components/pages/home.vue'
import TodoListTest from '../components/pages/TodoListTest.vue'
import TodoList from '../components/pages/TodoList.vue'
import profile from '../components/pages/profile.vue'

Vue.use(Router);

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
            component: TodoList
        },

        {
            path: '/product/list',
            name: 'Produkte',
            show_in_nav: true,
            component: TodoListTest
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
import Vue from 'vue/dist/vue.esm'
import Vuex from 'vuex'
import todos from './modules/todos'

Vue.use(Vuex);

export default new Vuex.Store({
    strict: process.env.NODE_ENV !== 'production',
    modules: {
        todos,
    },
})
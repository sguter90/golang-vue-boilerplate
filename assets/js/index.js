import "semantic-ui-css/semantic.min"
import Vue from 'vue/dist/vue.esm'
import VueRouter from 'vue-router'
import router from './components/router/index'
import App from './components/App.vue'

Vue.use(VueRouter);

new Vue({
    el: '#app',
    router,
    components: { App },
    template: '<App />',
});

jQuery('.ui.standard.dropdown').dropdown();
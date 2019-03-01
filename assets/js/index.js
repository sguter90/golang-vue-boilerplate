import "semantic-ui-css/semantic.min"
import Vue from 'vue/dist/vue.esm'
import router from './router/index'
import store from './store/index'
import App from './components/App.vue'

new Vue({
    el: '#app',
    store,
    router,
    render: h => h(App)
});

jQuery('.ui.standard.dropdown').dropdown();
import "semantic-ui-css/semantic.min"
import Vue from 'vue/dist/vue.esm'
import App from './components/App.vue'

let app = new Vue({
    el: '#app',
    render: h => h(App)
});

jQuery('.ui.standard.dropdown').dropdown();
<style scoped>

    #app-nav .item:hover {
        background: none #e9e9e9;
        color: rgba(0, 0, 0, 0.95);
        border-color: #d4d4d5;
    }

    .mobile.menu {
        display: block;
    }

    .desktop.menu {
        display: none;
    }

    @media only screen and (min-width: 992px) {
        .mobile.menu {
            display: none;
        }

        .desktop.menu {
            display: block;
        }
    }

</style>

<template>

    <div id="app-nav">

        <div class="ui vertical fluid tabular desktop menu">
            <router-link
                    active-class="active"
                    :class="{ active: route.path === $router.currentRoute.path } + ' item'"
                    v-for="route in $router.options.routes"
                    v-if="route.show_in_nav"
                    :key="'desktop-' + route.path"
                    :to="route.path"
            >
                {{ route.name }}
            </router-link>
        </div>


        <div class="ui sidebar vertical mobile menu">
            <router-link
                    active-class="active"
                    :class="{ active: route.path === $router.currentRoute.path } + ' item'"
                    v-for="route in $router.options.routes"
                    v-if="route.show_in_nav"
                    :key="'mobile-' + route.path"
                    :to="route.path"
            >
                {{ route.name }}
            </router-link>
        </div>

    </div>


</template>


<script>
    export default {
        name: 'nav',
        data () {
            return {
                msg: 'Welcome to Your Vue.js App'
            }
        }
    }
</script>
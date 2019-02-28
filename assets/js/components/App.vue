<template>
    <div id="app" class="ui centered grid pusher">
        <AppHeader></AppHeader>
        <div class="row">
            <div class="sixteen wide mobile  sixteen wide tablet fifteen wide computer column">
                <div class="ui grid">
                    <div class="two wide column computer only">
                        <app-nav></app-nav>
                    </div>
                    <div class="sixteen wide mobile  sixteen wide tablet  fourteen wide computer stretched column">
                        <div id="content" class="ui segment">
                            <router-view/>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import AppNav from './base/AppNav.vue'
    import AppHeader from './base/AppHeader.vue'

    export default {
        name: 'App',
        components: {AppNav, AppHeader}
    }
</script>
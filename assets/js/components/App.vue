<template>
    <div id="app" class="ui centered grid">
        <AppHeader></AppHeader>
        <div class="row">
            <div class="fifteen wide column">
                <div class="ui grid">
                    <div class="two wide column">
                        <app-nav></app-nav>
                    </div>
                    <div class="fourteen wide stretched column">
                        <div id="content" class="ui segment">
                            This is an stretched grid column. This segment will always match the tab height
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
        components: {AppNav, AppHeader}
    }
</script>
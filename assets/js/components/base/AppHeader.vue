<style scoped>
    .profile-switcher > .item,
    .profile-switcher > .item:hover {
        padding: 15px;
        z-index: 2000;
        border-radius: 5px;
        margin-left: 15px;
        background-color: #0F6640;
        display: inline-block;
    }

    .profile-switcher .item a {
        color: #fdfdfd;
        letter-spacing: 0.7px;
    }

    .profile-switcher .item i.icon {
        margin: 2px;
        color: #ffffff;
    }

    .logo img {
        max-height: 50px;
    }

</style>


<template>
    <div class="ui grid sixteen wide mobile  sixteen wide tablet fifteen wide computer column">
        <div class="row" style="border-bottom: 1px solid #d4d4d5">
            <div class="ui mobile tablet only four wide column vcenter">
                <a class="launch icon" onclick="jQuery('.ui.sidebar').sidebar('toggle');">
                    <i class="content icon"></i>
                </a>
            </div>
            <div class="ui eight wide logo column computer only">
                <router-link
                        to="/"
                >
                    <img src="static/images/logo.png" alt="logo"/>
                </router-link>
            </div>

            <div class="ui twelve wide mobile  twelve wide tablet  eight wide computer  profile right aligned column">
                <div class="right aligned menu profile-switcher">
                    <div class="item">
                        <div class="ui inline labeled icon top right pointing standard dropdown">
                            <a href="#">
                                <i class="ui rounded user icon"></i>
                                Username
                                <i class="chevron down icon"></i>
                            </a>
                            <div class="menu">
                                <router-link
                                        class="item"
                                        to="/profile"
                                >
                                    <i class="black setting icon"></i>
                                    Profil bearbeiten
                                </router-link>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>

    </div>


</template>
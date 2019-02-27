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

</style>


<template>
    <div class="ui grid fifteen wide column">
        <div class="row" style="border-bottom: 1px solid #d4d4d5">
            <div class="ui three wide column"></div>
            <div class="ui center aligned ten wide column">
                <img src="static/images/logo.png" style="max-height: 80px" alt="logo"/>
            </div>

            <div class="ui three wide right aligned column"
                 style="display: flex; justify-content: end; align-items: center;">
                <div class="right aligned menu profile-switcher">
                    <div class="item">


                        <i class="ui rounded user icon"></i>


                        <div class="ui inline labeled icon top right pointing standard dropdown">
                            <a href="#">
                                Username
                                <i class="chevron down icon"></i>
                            </a>
                            <div class="menu">
                                <a class="item" href="index.php?page=profile_edit">
                                    <i class="black setting icon"></i>
                                    Profil bearbeiten
                                </a>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>

    </div>


</template>
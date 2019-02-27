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
            <a class="item">
                Home
            </a>
            <a class="active item">
                Meine Aufgaben
            </a>
            <a class="item">
                Produkte
            </a>
            <a class="item">
                Zeitplanung
            </a>
        </div>


        <div class="ui sidebar vertical mobile menu">
            <a class="item">
                Home
            </a>
            <a class="active item">
                Meine Aufgaben
            </a>
            <a class="item">
                Produkte
            </a>
            <a class="item">
                Zeitplanung
            </a>
        </div>

    </div>


</template>
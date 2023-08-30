<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
    <RouterView />
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Lato&family=Montserrat:wght@300&family=Nunito+Sans:wght@500&display=swap');

*, *::before, *::after {
    box-sizing: border-box;
}

.container {
    max-width: 1440px;
    margin: 0 auto;
    padding: 0 1rem;
}

html {
    font-family: sans-serif;
}

body {
    margin: 0;
    padding: 0;
}

h1 {
    font-family: 'Montserrat', sans-serif;
}
</style>

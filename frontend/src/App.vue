<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
    <div class="container">
        <div class="flex-divide">
            <header>
                <h1>Cord Bleibaum</h1>
                <h2>making products from dreams</h2>
                <nav>
                    <RouterLink to="/">Home</RouterLink>
                    <RouterLink to="/projects">Projects</RouterLink>
                    <RouterLink to="/contact">Contact me</RouterLink>
                    <RouterLink to="/legal">Legal</RouterLink>
                </nav>
            </header>
            <RouterView />
        </div>
    </div>
</template>

<style>
@font-face {
  font-family: 'Montserrat';
  font-display: swap;
  src: url(@/assets/fonts/Montserrat.ttf) format('truetype');
}

@font-face {
  font-family: 'Open Sans';
  font-display: swap;
  src: url(@/assets/fonts/OpenSans.ttf) format('truetype');
}

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

h2, h3, h4 {
    font-family: 'Open Sans', sans-serif;
}

h1 {
    font-family: 'Montserrat', sans-serif;
    font-weight: 300;
    font-size: 64px;
    margin: 0;
    padding: 0;
}

h1::first-letter {
    color: #13b0ba;
}

.flex-divide {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    min-height: 100vh;
    align-items: center;
    gap: 64px;
}

.flex-divide > * {
    flex: 1;
}

nav {
    border-top: 2px solid #000;
    padding-top: 16px;
}

nav a {
    padding-right: 16px;
}

header {
    text-align: center;
}

a {
    text-decoration: none;
    color: #000;
    font-family: 'Open Sans', sans-serif;
}

p, li {
    font-family: 'Open Sans', sans-serif;
}

@media screen and (max-width: 1024px) {
    .flex-divide {
        flex-direction: column;
        min-height: inherit;
    }

    header {
        margin-top: 64px;
    }
}
</style>

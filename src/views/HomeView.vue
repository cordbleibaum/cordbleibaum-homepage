<script setup>
</script>

<template>
    <header>
        <div class="container">
            <h1>Cord Bleibaum</h1>
            <h2>making products from dreams</h2>
        </div>
    </header>
    <main>
        <div class="container">
            <p>
                Having a broad experience in software development and a passion for entrepreneurship,<br>
                I focus on building products from starting from the early stages of ideation to the final product.
            </p>
        </div>
    </main>
    <footer>
        <div class="container">
            <RouterLink to="/">Home</RouterLink>
            <RouterLink to="/legal">Legal</RouterLink>
        </div>
    </footer>
</template>

<style scoped>
h1 {
    font-size: 64px;
    margin: 0;
    padding: 0;
}

header {
    padding-top: 32px;
}

a {
    text-decoration: none;
    color: #000;
    padding-right: 16px;
}
</style>

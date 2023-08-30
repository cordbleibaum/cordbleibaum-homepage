<script setup>
</script>

<template>
    <div class="container">
        <div class="flex-divide">
            <header>
                <h1>Cord Bleibaum</h1>
                <h2>making products from dreams</h2>
                <div class="menu">
                    <RouterLink to="/">Home</RouterLink>
                    <RouterLink to="/">Projects</RouterLink>
                    <RouterLink to="/">Contact me</RouterLink>
                    <RouterLink to="/legal">Legal</RouterLink>
                </div>
            </header>
            <main>
                <h2>Hello,</h2>
                <p>
                    I'm Cord. Having a broad experience in software development and a passion for entrepreneurship,
                    I focus on building products from starting from the early stages of ideation to the final product.
                </p>
                <p>
                    I love to work with people who are passionate about their ideas and want to make them a reality.
                    Continuous learning and improvement are an important part of my life, making me a great fit for
                    fast-paced environments.
                </p>
                <p>
                    Having a background in Computer Science, I have worked with a large variety of technologies.
                    These include web development, mobile development, machine learning, computer graphics, FPGA development,
                    and many more.
                </p>
                <p>
                    I'm always on the lookout for new opportunities. If you have an idea you want to make a reality,
                    feel free to contact me.
                </p>
            </main>
        </div>
    </div>
</template>

<style scoped>
h1 {
    font-size: 64px;
    margin: 0;
    padding: 0;
}

header {
    text-align: center;
}

a {
    text-decoration: none;
    color: #000;
    padding-right: 16px;
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

.menu {
    border-top: 2px solid #000;
    padding-top: 16px;
}


</style>

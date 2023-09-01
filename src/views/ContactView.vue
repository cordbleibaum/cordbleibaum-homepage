<template>
    <main>
        <h2>Contact me</h2>
        <p>
            If you want to contact me, feel free to send me an email at
            <a href="mailto:cordbleibaum@gmail.com">cordbleibaum@gmail.com</a>.
        </p>
        <p>
            You can also directly contact me via the contact form below:
        </p>
        <form>
            <div>
                <label for="name">Name</label>
                <input type="text" id="name" name="name" placeholder="Your name" required>
            </div>
            <div>
                <label for="email">Email</label>
                <input type="email" id="email" name="email" placeholder="Your email" required>
            </div>
            <div>
                <label for="message">Message</label>
                <textarea id="message" name="message" placeholder="Your message" required></textarea>
            </div>
            <button type="submit">Send</button>
        </form>
    </main>
</template>

<style scoped>
form {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input[type="text"], input[type="email"], textarea {
    width: 100%;
    padding: 5px;
    border: 0;
    border-bottom: 1px solid #000;
}

input[type="text"]:focus, input[type="email"]:focus, textarea:focus {
    outline: none;
    border-bottom: 1px solid #000;
}

button {
    max-width: 200px;
    padding: 5px;
    border: 1px solid #000;
    background-color: #fff;
    cursor: pointer;
}
</style>
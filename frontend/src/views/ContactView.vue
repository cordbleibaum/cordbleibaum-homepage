<script setup>
import { ref } from 'vue'

const name = ref('')
const email = ref('')
const message = ref('')

async function submit(event) {
    if (event) {
        event.preventDefault()
    }

    const response = await fetch('https://api.cordbleibaum.de/contact', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name.value,
            email: email.value,
            message: message.value
        })
    })

    name.value = ''
    email.value = ''
    message.value = ''

    if (response.ok) {
        alert('Your message has been sent!')
    } else {
        alert('An error occurred while sending your message.')
    }
}
</script>

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
                <input type="text" id="name" name="name" placeholder="Your name" required v-model="name">
            </div>
            <div>
                <label for="email">Email</label>
                <input type="email" id="email" name="email" placeholder="Your email" required v-model="email">
            </div>
            <div>
                <label for="message">Message</label>
                <textarea id="message" name="message" placeholder="Your message" required v-model="message"></textarea>
            </div>
            <button @click="submit">Send</button>
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
    border-bottom: 1px solid var(--main-color);
}

input[type="text"]:focus, input[type="email"]:focus, textarea:focus {
    outline: none;
    border-bottom: 1px solid var(--main-color);
}

button {
    max-width: 200px;
    padding: 5px;
    border: 1px solid var(--main-color);
    background-color: var(--background-color);
    cursor: pointer;
}
</style>
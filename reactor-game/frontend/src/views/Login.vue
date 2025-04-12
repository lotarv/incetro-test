<template>
    <div class="login-container">
        <h2>Login to Reactors Game</h2>
        <button @click="loginWithTelegram" class="telegram-button">Login with Telegram</button>
    </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {useRouter} from "vue-router" 
import axios from 'axios'
import {API_BASE_URL} from "../config"

export default defineComponent({
    name: "Login",
    setup() {
        const router = useRouter()
        async function loginWithTelegram() {
            if (!window.Telegram?.WebApp) {
                alert("Please open this app in Telegram")
                return
            }
            const initData = window.Telegram.WebApp.initData
            console.log(initData)
            if (!initData) {
                alert("Unable to get Telegram user data")
                return
            }
            console.log('Sending initData:', initData) // Лог для отладки

            try {
                const response = await axios.post(`${API_BASE_URL}/auth/telegram`, {initData})
                const userID = response.data.user_id
                localStorage.setItem('userID', userID)
                alert("Login successfull! User ID: " + userID)
                router.push('/bonuses')

            } catch(error) {
                console.log("Login error:")
                alert(`Login failed: ${error}`)
            }
        }

        return {
            loginWithTelegram
        }
    }
})
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: hsl(0, 0%, 95%);
}

.telegram-button {
  padding: 15px 30px;
  font-size: 18px;
  background-color: #0088cc;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

.telegram-button:hover {
  background-color: #006699;
}
</style>
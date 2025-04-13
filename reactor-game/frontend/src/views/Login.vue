<template>
    <div class="login-container">
        <TelegramIcon class="telegram-icon"/>
        <button @click="loginWithTelegram" class="telegram-button">Login with Telegram</button>
    </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {useRouter} from "vue-router" 
import axios from 'axios'
import {API_BASE_URL} from "../config"
import TelegramIcon from "../components/icons/TelegramIcon.vue"

export default defineComponent({
    name: "Login",
    components: {
        TelegramIcon,
    },
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
  background-color: hsl(0, 0%, 95%);
  height: 720px;
  gap:25px;
}
.telegram-icon{
    width:200px;
}
.telegram-button {
  padding: 15px 30px;
  font-size: 18px;
  background-color: #28A8E8;
  color: white;
  border: none;
  border-radius: 16px;
  cursor: pointer;
  font-weight: 600;
}

.telegram-button:hover {
  background-color: #006699;
}
</style>
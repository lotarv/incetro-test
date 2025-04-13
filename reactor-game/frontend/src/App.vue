<template>
  <div class="app-container">
    <div v-if="error" class="error-container">
        <p class="error">{{ error }}</p>
        <button @click="retryLogin" :disabled="isLoading">
            {{ isLoading? "Retrying..." : "Try Again" }}
        </button>
    </div>
    <div v-else class="content">
      <router-view></router-view>
    </div>
    <nav class="tab-bar">
      <router-link class="tab-link" to="/bonuses">
        <span class="tab-icon"><ChipikIcon/></span>
        <span class="tab-text">Bonuses</span>
      </router-link>
      <router-link class="tab-link" to="/reactors">
        <span class="tab-icon"><LightningIcon/></span>
        <span class="tab-text">Reactors</span>
      </router-link>
      <router-link  class="tab-link" to="/top">
        <span class="tab-icon"><TopIcon/></span>
        <span class="tab-text">Top</span>
      </router-link>
    </nav>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import ChipikIcon from "./components/icons/СhipikIcon.vue"
import LightningIcon from './components/icons/LightningIcon.vue'
import TopIcon from './components/icons/TopIcon.vue'
import {useRouter} from "vue-router"
import axios from 'axios'
import {API_BASE_URL} from "./config"
export default defineComponent({
  name: 'App',
  components: {
    ChipikIcon,
    LightningIcon,
    TopIcon,
  },
  setup() {
    const router = useRouter();
    const isLoading = ref(false);
    const error = ref('');

    const loginWithTelegram = async () => {

      isLoading.value = true;
      error.value = '';

      try {
        if (!window.Telegram?.WebApp) {
          throw new Error('Please open this app in Telegram');
        }

        const initData = window.Telegram.WebApp.initData;
        if (!initData) {
          throw new Error('Unable to get Telegram user data');
        }

        console.log('Sending initData:', initData);

        const response = await axios.post(`${API_BASE_URL}/auth/telegram`, {initData})
        const userId = response.data.user_id.toString();
        console.log("userID: ", userId)
        localStorage.setItem('userID', userId);

        router.push('/bonuses');
      } catch (err: any) {
        error.value = err.message || 'Login failed';
        console.error('Login error:', err);
      } finally {
        isLoading.value = false;
      }
    };

    const retryLogin = () => {
      loginWithTelegram();
    };

    // Запускаем авторизацию при загрузке
    onMounted(() => {
        console.log("Autorise")
        loginWithTelegram();
    });

    return {
      isLoading,
      error,
      retryLogin,
    };
  },
});

</script>

<style scoped>
.app-container {
  max-width: 480px;
  margin:0 auto;
  font-family: 'Montserrat', sans-serif;
  min-height: 720px;
  display: flex;
  flex-direction: column;
  border-radius: 15px;
  overflow: hidden;
}

.content {
  flex:1;
  overflow-y: auto;
}

.tab-bar {
  bottom:0;
  color:gray;
  text-decoration: none;
  display: flex;
  max-width: 480px;
  width:100%;
  justify-content: space-around;
  padding-top:10px;
  background-color: hsl(0,0%,100%);
  box-shadow: 0 -4px 6px -1px rgba(0, 0, 0, 0.1);

}
.tab-link{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap:5px;
  user-select: none;
}
.tab-link.router-link-active {
  color:#DA2F20;
}
.tab-icon{
  font-size: 24px;
}
.tab-text{
  font-size: 12px;
}


</style>
<template>
    <div class="top-container">
        <div class="title">Leaderboard</div>
        <div v-if="usersLeaderboard.length > 0" class="leaderboard">
            <div v-for="(user, index) in usersLeaderboard" class="row">
                <div class="font-bold">{{ index + 1 }}</div>
                <div>{{ user.username }}</div>
                <div class="ml-auto">
                    <span class="text-[#8A8A8F]">{{ user.balance }}</span>
                    <span><BalanceIcon class="text-xl"/></span>
                </div>
            </div>
        </div>
        <p v-else>Loading leaderboard...</p>
    </div>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref} from 'vue'
import axios from 'axios'
import BalanceIcon from '../components/icons/BalanceIcon.vue'
import {API_BASE_URL} from "../config"
interface UserInfo {
    username: string
    balance: number
}

export default defineComponent({
    name:"Top",
    components: {
        BalanceIcon
    },
    setup() {
        const usersLeaderboard = ref<UserInfo[]>([])
        onMounted(async () => {
            try{
                const response = await axios.get(`${API_BASE_URL}/top`)
                usersLeaderboard.value = response.data.rating;
            } catch(error) {
                console.error("Error fetching users leaderboard: ", error)
            }

        })
        return {
            usersLeaderboard,
        }
    }

})
</script>

<style scoped>
.top-container{
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding:20px;
    background-color: hsl(0, 0%, 95%);
    color:black;
    height:720px;
}
.title{
    font-size: 22px;
    letter-spacing: 0.35px;
    font-weight: 600;
    text-align: center;
}

.leaderboard{
    display: flex;
    flex-direction: column;
    gap:10px;
}

.row{
    display: flex;
    align-items: center;
    padding:20px 15px;
    gap:15px;
    border-bottom: 0.5px solid #C6C6C8;
}

.row div:last-child{
    display: flex;
    gap:10px;
    align-items: center;
}
</style>
<template>
    <div class="bonuses-container">
        <div class="balance-item">
            <span class="bonuses-icon"><BalanceIcon/></span>
            <span>{{ bonusesStore.state.balance }}</span>
        </div>
        <div class="active-reactor-item">
            <div class="collected-bonuses">
                <span class="bonuses-icon" ><BalanceIcon/></span>
                <span>
                    {{ 
                        bonusesStore.state.farm_status === "claim"
                        ? bonusesStore.state.active_reactor.tokens_per_cycle
                        : Math.floor((bonusesStore.state.progress / 100) * bonusesStore.state.active_reactor.tokens_per_cycle)
                     }}
                </span>
            </div>
            <div class="reactor-image-container">
                <img class="reactor-image" src="https://s3-alpha-sig.figma.com/img/6a83/9f6a/d4c56bfbe7e651a752eec8a82a16ed3b?Expires=1745193600&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=LYZzPtrQ4kN0a2mWiMFpTSvtFyMuNL8FhFA5FljksbGQjgvNoi9JWq5l1uz0M2c8gb7i5RR9KMxI~KEj1VJHsmdbYsjWo8WSB94uHRXxcJ4uEEZb-sfKTkWyAQXpBbAPBkJono4k28i1HzbMCZQh4~FEl4fOX6~MjSlmWEaCT9oyPz43nqSj1RbH9ZwKKvb-qAfEeKdmxiIwHPyEYExAlZR9EHw1xGvadxpjscrxwCQyhUci74Q0bmpGbXctAP-GOdK5deC2uIyC6QG9ji4PpM6lweciVvOGpO6Hm7kb7mCqcN7E2nDC8gU8Bo8-lAOjU-LjTBCuxIvZs28mSFbb4Q__" alt="">
            </div>
        </div>
        <button
         class="main-button"
         :disabled="bonusesStore.state.farm_status=='farming'"
         @click="handleButtonClick">
         <span v-if="bonusesStore.state.farm_status === 'start'">Start</span>
         <span v-else-if="bonusesStore.state.farm_status === 'farming'">Farming ({{ bonusesStore.state.time_left }}s)</span>
         <span v-else-if="bonusesStore.state.farm_status==='claim'">
        Claim {{ bonusesStore.state.active_reactor.tokens_per_cycle }}</span>
        </button>
    </div>
</template>


<script lang="ts">
import {defineComponent, onMounted, onUnmounted} from 'vue'
import { useBonusesStore } from '@/stores/bonuses';
import BalanceIcon from '../components/icons/BalanceIcon.vue'
export default defineComponent({
    name:"Bonuses",
    components: {
        BalanceIcon,
    },

    setup() {
        const bonusesStore = useBonusesStore()

        onMounted(() => {
            bonusesStore.fetchBonuses()
        })

        onUnmounted(() => {
            bonusesStore.clearTimer()
        })

        function handleButtonClick() {
            if (bonusesStore.state.farm_status === "start") {
                bonusesStore.startFarming()
            } else if (bonusesStore.state.farm_status === "claim") {
                bonusesStore.claimBonuses()
            }
        }

        return {
            bonusesStore,
            handleButtonClick
        }
    }
})
</script>

<style scoped>
.bonuses-container{
    background-color: hsl(0, 0%, 95%);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height:720px;
    padding:15px;
    color:black;
}

.balance-item{
    padding:20px 20px;
    text-align: center;
    background-color: white;
    border-radius: 16px;
    display: flex;
    justify-content: center;
    font-weight: 600;
    gap:5px;
    margin-top:20px;
    font-size:20px;
    line-height: 28px;
    letter-spacing: 0.35px;
}

.bonuses-icon{
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.collected-bonuses{
    font-size: 42px;
    display: flex;
    justify-content: center;
    font-weight: 600;
    gap: 8px;
}
.collected-bonuses .bonuses-icon{
    font-size: 52px;
}

.active-reactor-item{
    display: flex;
    flex-direction: column;
    gap:30px;
}

.reactor-image-container{
    display: flex;
    justify-content: center;
    align-items: center;
}

.reactor-image{
    border-radius: 50%;
    width:250px;
    height:250px;
}

.main-button{
    padding:15px 10px;
    text-align: center;
    background-color: #DA2F20;
    border-radius: 13px;
    display: flex;
    justify-content: center;
    font-weight: 600;
    gap:5px;
    margin-top:20px;
    font-size:17px;
    line-height: 28px;
    letter-spacing: 0.35px;
    color:white;
}
</style>
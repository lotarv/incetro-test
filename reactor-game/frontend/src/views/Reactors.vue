<template>
    <div class="reactors-container">
        <h2>Reactors</h2>
        <div v-if="isLoading" class="loading">Loading reactors..</div>
        <Carousel v-else-if="reactorsStore.state.reactors.length > 0" :value="reactorsStore.state.reactors" :numVisible="1" :numScroll="1" :showIndicators="true"
            :showNavigators="false" :circular="false" v-model:page="activeIndex" class="reactors-carousel">
            <template #item="slotProps">
                <div class="slide">
                    <img class="reactor-image" :src="`https://s3-alpha-sig.figma.com/img/6a83/9f6a/d4c56bfbe7e651a752eec8a82a16ed3b?Expires=1745193600&Key-Pair-Id=APKAQ4GOSFWCW27IBOMQ&Signature=LYZzPtrQ4kN0a2mWiMFpTSvtFyMuNL8FhFA5FljksbGQjgvNoi9JWq5l1uz0M2c8gb7i5RR9KMxI~KEj1VJHsmdbYsjWo8WSB94uHRXxcJ4uEEZb-sfKTkWyAQXpBbAPBkJono4k28i1HzbMCZQh4~FEl4fOX6~MjSlmWEaCT9oyPz43nqSj1RbH9ZwKKvb-qAfEeKdmxiIwHPyEYExAlZR9EHw1xGvadxpjscrxwCQyhUci74Q0bmpGbXctAP-GOdK5deC2uIyC6QG9ji4PpM6lweciVvOGpO6Hm7kb7mCqcN7E2nDC8gU8Bo8-lAOjU-LjTBCuxIvZs28mSFbb4Q__`"
                        alt="Reactor" />
                </div>
            </template>
        </Carousel>
        <div v-else class="no-reactors">No reactors available</div>
        <div v-if="reactorsStore.state.reactors.length > 0" class="reactor-info">
            <div>
                <span>Farm Time</span>
                <span>{{ reactorsStore.state.reactors[activeIndex].farmTime }} s</span>
            </div>
            <div>
                <span>Tokens per Cycle</span>
                <p class="tokens_value">
                    <BalanceIcon style="font-size: 22px;" />
                    <span>{{ reactorsStore.state.reactors[activeIndex].tokensPerCycle }}</span>
                </p>
            </div>
        </div>

        <button
        :class="{'action-button':true,
            'active':isButtonActive(),
            'inactive': !isButtonActive(),
            'owned': isReactorOwned(),
            'buy': !isReactorOwned(),
            'use': isReactorOwned() && userStore.state.ActiveReactor !== reactorsStore.state.reactors[activeIndex].id
        }"
        :disabled="!isButtonActive()"
        @click="handleButtonClick">
        <span v-if="reactorsStore.state.reactors.length > 0 && userStore.state.ActiveReactor === reactorsStore.state.reactors[activeIndex].id">In use</span>
        <div v-else-if="reactorsStore.state.reactors.length > 0 && !isReactorOwned()">
            <p class="tokens_value">
                    <span>Buy for</span>
                    <BalanceIcon class="text-black text-[22px]" />
                    <span>{{ reactorsStore.state.reactors[activeIndex].price }}</span>
            </p>
        </div>
        <span v-else>Use</span>
        </button>
    </div>
</template>
  
<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { useReactorsStore } from '@/stores/reactors'
import { useUserStore } from '@/stores/user'
import BalanceIcon from '@/components/icons/BalanceIcon.vue'
import Carousel from 'primevue/carousel'

export default defineComponent({
    name: 'Reactors',
    components: {
        Carousel,
        BalanceIcon
    },
    setup() {
        const reactorsStore = useReactorsStore()
        const activeIndex = ref(0)

        const userStore = useUserStore()
        const isLoading = ref(true) 

        onMounted(async () => {
            try{
                await Promise.all([reactorsStore.fetchReactors(), userStore.fetchUser()])
                console.log("Reactors loaded")
                console.log("User loaded: ", userStore.state.Reactors)
            } catch(error) {
                console.error("Failed to load data: ", error)
            } finally {
                isLoading.value = false
            }
        })

        function isReactorOwned() {
            return userStore.state.Reactors.some(r => r.id === reactorsStore.state.reactors[activeIndex.value]?.id)
        }

        function isButtonActive() {
            const currentReactor = reactorsStore.state.reactors[activeIndex.value]
            if (!currentReactor) return false
            if (userStore.state.ActiveReactor === currentReactor.id) return false
            if (!isReactorOwned()) return userStore.state.Balance >= currentReactor.price
            return userStore.state.FarmStatus !== 'farming'
        }

        async function handleButtonClick() {
            const currentReactor = reactorsStore.state.reactors[activeIndex.value]
            if (!currentReactor) return

            if (userStore.state.ActiveReactor === currentReactor.id) return

            try {
                if (!isReactorOwned()) {
                    await userStore.buyReactor(currentReactor.id)
                    alert("Reactor has been bought!")
                } else {
                    await userStore.useReactor(currentReactor.id)
                } 
            } catch(error) {
                console.error("Action failed: ", error)
            }
        }

        return {
            reactorsStore,
            activeIndex,
            userStore,
            isReactorOwned,
            isButtonActive,
            handleButtonClick,
            isLoading
        }
    }
})
</script>
  
<style scoped>
.reactors-container {
    padding: 15px;
    background-color: hsl(0, 0%, 100%);
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap:10px;
}

h2 {
    margin-bottom: 20px;
    font-size: 24px;
    font-weight: 600;
}

.reactors-carousel {
    width: 100%;
}

.slide {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 10px;
}

.reactor-image {
    border-radius: 50%;
    width:250px;
    height:250px;
}

.reactor-info {
    text-align: center;
    width: 100%;
    background-color: #F2F2F7;
    border-radius: 15px;
}

.tokens_value {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 5px;
}

.reactor-info div {
    margin: 5px 0;
    font-size: 16px;
    display: flex;
    padding: 10px 20px;
    justify-content: space-between;
}

.reactor-info div:last-child {
    border-top: 0.5px solid #C6C6C8;
}

.reactors-carousel :deep(.p-carousel-viewport) {
    overflow-x: hidden;
}

.reactors-carousel :deep(.p-carousel-indicators) {
    padding: 10px 0;
}

/* Стили для всех индикаторов */
.reactors-carousel :deep(.p-carousel-indicator .p-carousel-indicator-button) {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background-color: #ccc;
    border: none;
    margin: 0 6px;
    transition: all 0.3s;
}

/* Стили для активного индикатора */
.reactors-carousel :deep(.p-carousel-indicator-active .p-carousel-indicator-button) {
    background-color: #DA2F20;
    width: 10px;
    height: 10px;
}

.action-button{
    background-color: #DA2F20;
    width:100%;
    border-radius: 13px;
    padding:16px;
    color:white;
    font-weight: bold;
    font-size: 17px;
    margin-top: 20px;
}

.inactive {
    background-color: #DA2F209E;
}


</style>
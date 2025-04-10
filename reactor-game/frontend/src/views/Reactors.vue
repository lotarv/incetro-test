<template>
    <div class="reactors-container">
        <h2>Reactors</h2>
        <Carousel :value="reactorsStore.state.reactors" :numVisible="1" :numScroll="1" :showIndicators="true"
            :showNavigators="false" :circular="false" v-model:page="activeIndex" class="reactors-carousel">
            <template #item="slotProps">
                <div class="slide">
                    <img class="reactor-image" :src="`https://placehold.co/400x400?text=Reactor+${slotProps.data.id}`"
                        alt="Reactor" />
                </div>
            </template>
        </Carousel>
        <div v-if="reactorsStore.state.reactors.length > 0" class="reactor-info">
            <div>
                <span>Farm Time</span>
                <span>{{ reactorsStore.state.reactors[activeIndex].farmTime }} с</span>
            </div>
            <div>
                <span>Tokens per Cycle</span>
                <p class="tokens_value">
                    <BalanceIcon style="font-size: 22px;"/>
                    <span>{{ reactorsStore.state.reactors[activeIndex].tokensPerCycle }}</span>
                </p>
            </div>
        </div>

        <div class="buy-select-button">
            
        </div>
    </div>
</template>
  
<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useReactorsStore } from '@/stores/reactors'
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

        onMounted(() => {
            reactorsStore.fetchReactors()
        })

        return {
            reactorsStore,
            activeIndex,
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
}

.reactor-info {
    margin-top: 10px;
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
</style>
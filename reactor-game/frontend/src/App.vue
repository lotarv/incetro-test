<template>
  <div class="bonuses">
    <h1>Bonuses</h1>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <p>Balance: {{ bonuses.balance }}</p>
      <p>Active Reactor: {{ bonuses.active_reactor.id }} ({{ bonuses.active_reactor.farm_time }}s, {{ bonuses.active_reactor.tokens_per_cycle }} tokens)</p>
      <p>Status: {{ bonuses.farm_status }}</p>
      <p>Progress: {{ bonuses.progress }}%</p>
      <p>Time Left: {{ bonuses.time_left }}s</p>
      <button @click="startFarming" :disabled="bonuses.farm_status !== 'start'">Start Farming</button>
      <button @click="claimBonuses" :disabled="bonuses.farm_status !== 'completed'">Claim Bonuses</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

// Типы данных
interface Reactor {
  id: number;
  farm_time: number;
  tokens_per_cycle: number;
  price: number;
}

interface Bonuses {
  balance: number;
  active_reactor: Reactor;
  farm_status: string;
  progress: number;
  time_left: number;
}

export default defineComponent({
  name: 'Bonuses',
  data(): { bonuses: Bonuses; loading: boolean } {
    return {
      bonuses: {
        balance: 0,
        active_reactor: { id: 0, farm_time: 0, tokens_per_cycle: 0, price: 0 },
        farm_status: 'start',
        progress: 0,
        time_left: 0,
      },
      loading: true,
    };
  },
  methods: {
    async fetchBonuses() {
      try {
        const response = await axios.get<Bonuses>('http://localhost:8080/bonuses');
        this.bonuses = response.data;
        this.loading = false;
      } catch (error) {
        console.error('Error fetching bonuses:', error);
      }
    },
    async startFarming() {
      try {
        await axios.post('http://localhost:8080/bonuses/start');
        await this.fetchBonuses();
      } catch (error) {
        console.error('Error starting farming:', error);
      }
    },
    async claimBonuses() {
      try {
        await axios.post('http://localhost:8080/bonuses/claim');
        await this.fetchBonuses();
      } catch (error) {
        console.error('Error claiming bonuses:', error);
      }
    },
  },
  created() {
    this.fetchBonuses();
  },
});
</script>

<style scoped>
.bonuses {
  text-align: center;
  padding: 20px;
}
button {
  margin: 5px;
  padding: 10px;
  cursor: pointer;
}
button:disabled {
  cursor: not-allowed;
  opacity: 0.5;
}
</style>
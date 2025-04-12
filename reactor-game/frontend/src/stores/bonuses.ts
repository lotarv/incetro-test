import { defineStore } from 'pinia'
import axios from 'axios'
import { reactive } from 'vue'
import {API_BASE_URL} from "../config"
// Интерфейс для реактора
interface Reactor {
  id: number
  farm_time: number
  tokens_per_cycle: number
  price: number
}

// Интерфейс для состояния
interface BonusState {
  balance: number
  active_reactor: Reactor
  farm_status: 'start' | 'farming' | 'claim'
  progress: number
  time_left: number
}

// Создаём стор
export const useBonusesStore = defineStore('bonuses', () => {
  // Реактивное состояние
  const state = reactive<BonusState>({
    balance: 0,
    active_reactor: { id: 0, farm_time: 0, tokens_per_cycle: 0, price: 0 },
    farm_status: 'start',
    progress: 0,
    time_left: 0
  })

  // Локальная переменная для таймера
  let timer: number = 0

  // Действия
  const fetchBonuses = async () => {
    try {
      const response = await axios.get<BonusState>(`${API_BASE_URL}/bonuses?userID=${localStorage.getItem('userID')}`)
      Object.assign(state, response.data)
      console.log(state)
      if (state.farm_status === 'farming') {
        startTimer()
      }
    } catch (err) {
      console.error('Failed to fetch bonuses:', err)
    }
  }

  const startFarming = async () => {
    if (state.farm_status !== 'start') return
    try {
      await axios.post(`${API_BASE_URL}/bonuses/start?userID=${localStorage.getItem('userID')}`)
      await fetchBonuses()
    } catch (err) {
      console.error('Failed to start farming:', err)
    }
  }

  const claimBonuses = async () => {
    if (state.farm_status !== 'claim') return
    try {
      await axios.post(`${API_BASE_URL}/bonuses/claim?userID=${localStorage.getItem('userID')}`)
      await fetchBonuses()
    } catch (error) {
      console.error('Failed to claim bonuses:', error)
    }
  }

  const startTimer = () => {
    if (timer) clearInterval(timer) 
    timer = setInterval(() => {
      if (state.time_left > 0) {
        state.time_left -= 1
        state.progress = Math.min(
          ((state.active_reactor.farm_time - state.time_left) / state.active_reactor.farm_time) * 100,
          100
        ) 
      }
      if (state.time_left <= 0 && state.farm_status === 'farming') {
        state.farm_status = 'claim'
        state.progress = 100
        clearInterval(timer) 
      }
    }, 1000)
  }

  const clearTimer = () => {
    if (timer) clearInterval(timer) 
  }

  // Возвращаем состояние и методы
  return {
    state,
    fetchBonuses,
    startFarming,
    claimBonuses,
    startTimer,
    clearTimer
  }
})
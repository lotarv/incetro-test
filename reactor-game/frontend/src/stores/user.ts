import { defineStore } from "pinia";
import axios from "axios";

interface Reactor {
    id: number
    farm_time: number
    tokens_per_cycle: number
    price:number
}

interface UserState {
    balance: number,
    active_reactor: Reactor,
    farm_status: string,
    progress: number,
    time_left:number
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        balance:0,
        active_reactor: {id: 0, farm_time:0, tokens_per_cycle: 0, price: 0},
        farm_status:'start',
        progress:0,
        time_left:0
    }),
    actions: {
        async fetchBonuses() {
            try {
                const response = await axios.get<UserState>('http://localhost:8080/bonuses')
                Object.assign(this, response.data);
            } catch(error) {
                console.error("Failed to fetch bonuses: ", error)
            }
        },

        async StartFarming() {
            try {
                await axios.post('http://localhost:8080/bonuses/start')
                await this.fetchBonuses()
            } catch(error) {
                console.error("Failed to start farming: ", error)
            }
        },

        async claimBonuses() {
            try {
                await axios.post('http://localhost:8080/bonuses/claim')
                await this.fetchBonuses()
            } catch(error) {
                console.error("Failed to claim bonuses: ", error)
            }
        }
    }
})
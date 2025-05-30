import { defineStore } from "pinia";
import axios from "axios";
import { reactive } from "vue";
import {API_BASE_URL} from "../config"
interface Reactor {
    id:number
    farmTime: number
    tokensPerCycle: number
    price: number
}

interface ReactorState {
    reactors: Reactor[]
}

export const useReactorsStore = defineStore('reactors', () => {
    const state = reactive<ReactorState>({
        reactors: [],
    })

    async function fetchReactors() {
        try {
            const response = await axios.get(`${API_BASE_URL}/reactors?userID=${localStorage.getItem('userID')}`)
            state.reactors = response.data.map((r: any) => ({
                id: r.id,
                farmTime: r.farm_time,
                tokensPerCycle: r.tokens_per_cycle,
                price: r.price
            }))
        } catch(error) {
            console.error("Failed to fetch reactors: ", error)
        }
    }

    return {
        state,
        fetchReactors,
    }
})
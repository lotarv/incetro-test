import { defineStore } from "pinia";
import { reactive } from "vue";
import axios from "axios";



interface Reactor {
    id:number
    farmTime: number
    tokensPerCycle: number
    price: number
}

interface UserState {
    ID: number
    Balance: number
    ActiveReactor: number
    FarmStatus: 'start' | 'farming' | 'claim'
    Reactors: Reactor[]
}

export const useUserStore = defineStore('user', () => {
    const state = reactive<UserState>({
        ID: 0,
        Balance: 0,
        ActiveReactor: 0,
        FarmStatus: 'start',
        Reactors: [] as Reactor[]
    })

    async function fetchUser() {
        try {
            const response = await axios.get("http://localhost:8080/user")
            state.ID = response.data.id;
            state.Balance = response.data.balance;
            state.ActiveReactor = response.data.active_reactor;
            state.FarmStatus = response.data.farm_status
            state.Reactors = response.data.reactors;
        } catch(error) {
            console.error("Failed to fetch user: ", error)
        }
    }

    async function buyReactor(reactorID: number) {
        try {
            await axios.post(`http://localhost:8080/reactors/buy/${reactorID}`)
            await fetchUser()
            return true
        } catch(error) {
            console.error("Failed to buy reactor: ", error)
            throw error
        }
    }

    async function useReactor(reactorID: number) {
        try {
            await axios.post(`http://localhost:8080/reactors/use/${reactorID}`)
            await fetchUser()
            return true
        } catch(error) {
            console.error("Failed to use reactor: ", error)
            throw error
        }
    }

    return {
        state,
        fetchUser,
        buyReactor,
        useReactor
    }


})
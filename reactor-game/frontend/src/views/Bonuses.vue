<template>
    <div class="bonuses-container">
        <div class="balance-item">
            <span class="bonuses-icon"><BalanceIcon/></span>
            <span>1000</span>
        </div>
        <div class="active-reactor-item">
            <div class="collected-bonuses">
                <span class="bonuses-icon" ><BalanceIcon/></span>
                <span>0</span>
            </div>
            <div class="reactor-image-container">
                <img class="reactor-image" src="https://placehold.co/350x350" alt="">
            </div>
        </div>
        <button class="main-button">
            Start
        </button>
    </div>
</template>


<script lang="ts">
import {defineComponent} from 'vue'
import BalanceIcon from '../components/icons/BalanceIcon.vue'
export default defineComponent({
    name:"Bonuses",
    components: {
        BalanceIcon,
    }
})
</script>

<style scoped>
.bonuses-container{
    background-color: hsl(0, 0%, 95%);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height:100%;
    padding:15px;
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
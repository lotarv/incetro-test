<template>
  <div class="app-container">
    <div class="content">
      <router-view></router-view>
    </div>
    <nav v-if="localStorage.getItem('userID')" class="tab-bar">
      <router-link class="tab-link" to="/bonuses">
        <span class="tab-icon"><ChipikIcon/></span>
        <span class="tab-text">Bonuses</span>
      </router-link>
      <router-link class="tab-link" to="/reactors">
        <span class="tab-icon"><LightningIcon/></span>
        <span class="tab-text">Reactors</span>
      </router-link>
      <router-link  class="tab-link" to="/top">
        <span class="tab-icon"><TopIcon/></span>
        <span class="tab-text">Top</span>
      </router-link>
    </nav>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ChipikIcon from "./components/icons/СhipikIcon.vue"
import LightningIcon from './components/icons/LightningIcon.vue'
import TopIcon from './components/icons/TopIcon.vue'
export default defineComponent({
  name: 'App',
  components: {
    ChipikIcon,
    LightningIcon,
    TopIcon,
  },
  setup() {
    return {
      localStorage
    }
  }
})

</script>

<style scoped>
.app-container {
  max-width: 480px;
  margin:0 auto;
  font-family: 'Montserrat', sans-serif;
  border: 1px solid #000;
  height:100vh;
  max-height: 930px;
  display: flex;
  flex-direction: column;
  border-radius: 15px;
  overflow: hidden;
}

.content {
  flex:1;
  overflow-y: auto;
}

.tab-bar {
  bottom:0;
  color:gray;
  text-decoration: none;
  display: flex;
  max-width: 480px;
  width:100%;
  justify-content: space-around;
  padding-top:10px;

}
.tab-link{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap:5px;
  user-select: none;
}
.tab-link.router-link-active {
  color:#DA2F20;
}
.tab-icon{
  font-size: 24px;
}
.tab-text{
  font-size: 12px;
}


</style>
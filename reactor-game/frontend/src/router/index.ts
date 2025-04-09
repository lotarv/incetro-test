import { createRouter, createWebHistory } from "vue-router";
import Bonuses from '../views/Bonuses.vue'
import Reactors from "../views/Reactors.vue"
import Top from "../views/Top.vue"

const routes = [
    {path: "/", redirect: "/bonuses"},
    {path: "/bonuses", component: Bonuses},
    {path: "/reactors", component: Reactors},
    {path: "/top", component: Top},
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
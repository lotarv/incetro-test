import { createRouter, createWebHistory } from "vue-router";
import Bonuses from '../views/Bonuses.vue'
import Reactors from "../views/Reactors.vue"
import Login from "../views/Login.vue"
import Top from "../views/Top.vue"

const routes = [
    {path: "/", redirect: "/login"},
    {path: "/bonuses", component: Bonuses},
    {path: "/reactors", component: Reactors},
    {path: "/top", component: Top},
    {path: "/login", component: Login}
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
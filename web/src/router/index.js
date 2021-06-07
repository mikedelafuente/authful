import { createWebHistory, createRouter } from "vue-router";
import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import ResetPassword from "@/views/ResetPassword.vue";
import Login from "@/views/Login.vue";
import Logout from "@/views/Logout.vue";
import AccountRegister from "@/views/AccountRegister.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/about",
    name: "about",
    component: About,
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/logout",
    name: "logout",
    component: Logout,
  },
  {
    path: "/account/register",
    name: "account_register",
    component: AccountRegister,
  },
  {
    path: "/reset-password",
    name: "account_reset",
    component: ResetPassword,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
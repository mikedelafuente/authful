import { createWebHistory, createRouter } from "vue-router";
import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import ResetPassword from "@/views/ResetPassword.vue";
import Login from "@/views/Login.vue";
import AccountRegister from "@/views/AccountRegister.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/account/register",
    name: "Register for an Account",
    component: AccountRegister,
  },
  {
    path: "/reset-password",
    name: "Reset Password",
    component: ResetPassword,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
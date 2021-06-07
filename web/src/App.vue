<template>
  <div>
    <header class="navbar navbar-expand-md navbar-dark bd-navbar">
      <nav class="container-xxl flex-wrap flex-md-nowrap" aria-label="Main navigation">
        <div class="collapse navbar-collapse" id="bdNavbar">
          <ul class="navbar-nav flex-row flex-wrap bd-navbar-nav pt-2 py-md-0">
            <li class="nav-item p-1">
              <router-link to="/" class="navbar-brand p-0 me-2">Authful</router-link>
            </li>
            <li class="nav-item col-6 col-md-auto">
              <router-link class="nav-link p-2" active-class="active" to="/">Home</router-link>
            </li>
            <li class="nav-item col-6 col-md-auto">
              <router-link class="nav-link p-2" active-class="active" to="/about">About</router-link>
            </li>
          </ul>
          <hr class="d-md-none text-white-50" />
          <ul class="navbar-nav flex-row flex-wrap ms-md-auto">
            
          </ul>
          <LoginLogoutButtons ref="loginLogoutButtons" :key="loginState" />
        </div>
      </nav>
    </header>
    <router-view />
  </div>
</template>

<script>
import { EventBus } from './event-bus';
import LoginLogoutButtons from "./components/LoginLogoutButtons.vue";
export default {
  name: "App",
  components: {
    LoginLogoutButtons,
  },
  data() {
    return {
      loginState: 0 
    }
  },
  mounted() {
    EventBus.on('login', this.handleLoginLogout)
    EventBus.on('logout', this.handleLoginLogout)
  },
  methods: {
    handleLoginLogout: function() {
      this.loginState += 1 // By incrementing the counter, the component is re-evaluated
    }
  },
  
};
</script>

<style>
.bd-navbar {
  padding: 0.75rem 0;
  background-color: #7952b3;
}
.bd-navbar .navbar-toggler {
  padding: 0;
  border: 0;
}
.bd-navbar .navbar-nav .nav-link {
  padding-right: 0.25rem;
  padding-left: 0.25rem;
  color: rgba(255, 255, 255, 0.85);
}
.bd-navbar .navbar-nav .nav-link:hover,
.bd-navbar .navbar-nav .nav-link:focus {
  color: #fff;
}
.bd-navbar .navbar-nav .nav-link.active {
  font-weight: 600;
  color: #fff;
}
.bd-navbar .navbar-nav-svg {
  width: 1rem;
  height: 1rem;
}
.yellow-outline,
.btn-bd-login {
  font-weight: 600;
  color: #ffe484;
  border-color: #ffe484;
}

.active.btn-bd-login,
.btn-bd-login:hover,
.btn-bd-login:active {
  color: #2a2730;
  background-color: #ffe484;
  border-color: #ffe484;
}
.btn-bd-login:focus {
  box-shadow: 0 0 0 3px rgba(255, 228, 132, 0.25);
}

@import url("~bootstrap/dist/css/bootstrap.min.css");
</style>

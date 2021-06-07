<template>
  <div>
    <router-link
      v-if="!isValidJwt"
      class="btn btn-bd-login d-lg-inline-block my-2 my-md-0 ms-md-3"
      active-class="active"
      to="/login"
      >Login</router-link
    >
    <ul v-else class="navbar-nav flex-row flex-wrap ms-md-auto">
      <li class="nav-item dropdown">
        <a
          class="nav-link "
          data-bs-toggle="dropdown"
          href="#"
          role="button"
          aria-expanded="false"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="currentColor"
            class="bi bi-person-circle"
            viewBox="0 0 16 16"
          >
            <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"></path>
            <path
              fill-rule="evenodd"
              d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"
            ></path>
          </svg>

          <small class="d-md-none ms-2">Profile</small></a
        >
        <ul class="dropdown-menu">
          <li><router-link
              class="dropdown-item"
              id="account-profile"
              :to="{ name: 'account_profile' }"
              >Profile</router-link
            ></li>
         <li><hr class="dropdown-divider" /></li>
          <li>
            <router-link
              class="dropdown-item"
              id="logout"
              :to="{name: 'logout'}"
              >Logout</router-link
            >
          </li>
        </ul>
      </li>
      <li class="nav-item col-6 col-md-auto"></li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  computed: {
    isValidJwt: function () {
      return this.validateJwt();
    },
  },
  methods: {
    validateJwt: function () {
      var currentJwt = this.$cookies.get("authfulJwt");
      if (currentJwt == null || typeof currentJwt == undefined) {
        return false;
      } else {
        var parsedJwt = this.parseJwt(currentJwt);
        if (parsedJwt == null || typeof parsedJwt == undefined) {
          return false;
        }

        return true;
      }
    },
    parseJwt: function (token) {
      var base64Url = token.split(".")[1];
      var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
      var jsonPayload = decodeURIComponent(
        atob(base64)
          .split("")
          .map(function (c) {
            return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
          })
          .join("")
      );

      var parsedJson = JSON.parse(jsonPayload);
      return parsedJson;
    },
  },
};
</script>

<style>
</style>
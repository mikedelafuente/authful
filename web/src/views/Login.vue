<template>
  <form class="form-signin my-3 px-3  align-items-center rounded-3 border shadow-lg">
    <img
      class="mb-4"
      src="../assets/WeekendClip.png"
      alt=""
      width="120"
      height="120"
    />
    <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

    <div class="form-floating">
      <input
        type="email"
        class="form-control"
        id="floatingInput"
        v-model="username"
        placeholder="name@example.com"
      />
      <label for="floatingInput">Email address</label>
    </div>
    <div class="form-floating">
      <input
        type="password"
        class="form-control"
        id="floatingPassword"
        v-model="password"
        placeholder="Password"
      />
      <label for="floatingPassword">Password</label>
    </div>

    <div class="checkbox mb-3">
      <label> <input type="checkbox" value="remember-me" /> Remember me </label>
    </div>
    <button
      class="w-100 btn btn-lg btn-primary"
      type="button"
      v-on:click="doLogin"
    >
      Sign in
    </button>
    <div class="mt-2 mb-3">
      <router-link to="/forgot-password" class="link-secondary">Forgot Password?</router-link>
    </div>
    <p class="mt-4 mb-3 text-muted">&copy; 2021</p>
  </form>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  props: {},
  methods: {
    doLogin: function () {
      const axios = require("axios").default;

      axios
        .post(
          "http://localhost:8081/api/v1/signin",
          {
            username: this.username,
            password: this.password,
          },
          {
            headers: {
              Accept: "application/json",
              "Content-Type": "application/json",
              Cache: "no-cache",
            },
            withCredentials: true,
          }
        )
        .then((response) => {
          this.$cookies.set(
            "userJwt",
            response.data.jwt,
            new Date(response.data.expires)
          );
          //this.response = JSON.stringify(response, null, "")
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style>
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-floating:focus-within {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
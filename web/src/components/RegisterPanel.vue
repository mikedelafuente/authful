<template>
  <div class="container-fluid">
    <div class="row justify-content-center">
      <form class="col-5 my-3 px-3 align-items-center rounded-3 border shadow-lg">
        <img class="mb-4" src="../assets/WeekendClip.png" alt width="120" height="120" />
        <h1 class="h3 mb-3 fw-normal">Register for an Account</h1>
        <div class="form-group py-3">
          <div class="form-floating">
            <input
              type="email"
              class="form-control"
              id="floatingInput"
              v-model="username"
              placeholder="name@example.com"
            />
            <label for="floatingInput" >Email address</label>
          </div>
        </div>
        <div class="form-group">
          <div class="form-floating py-1">
            <input
              type="password"
              class="form-control"
              id="floatingPassword"
              v-model="password"
              placeholder="Password"
            />
            <label for="floatingPassword">Password</label>
          </div>
          <div class="form-floating py-1">
            <input
              type="password"
              class="form-control"
              id="floatingConfirmPassword"
              v-model="confirmPassword"
              placeholder="Re-enter password"
            />
            <label for="floatingConfirmPassword">Re-enter password</label>
          </div>
        </div>
        <div class="form-group py-3">
        <button class="w-100 btn btn-lg btn-primary" type="button" v-on:click="doRegister">Register</button>
        <div class="mt-2 mb-3">
          Already have an account?
          <router-link to="/login" class="link-secondary">Login</router-link>
        </div>

        </div>
        <p class="mt-4 mb-3 text-muted">&copy; 2021</p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginPanel",
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
    };
  },
  props: {},
  methods: {
    doRegister: function () {
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

</style>
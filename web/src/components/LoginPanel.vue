<template>
  <form class="form-signin my-3 px-3 align-items-center rounded-3 border shadow-lg">
    <img class="mb-4" src="../assets/WeekendClip.png" alt width="120" height="120" />
    <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

    <div class="form-floating">
      <input
        type="email"
        class="form-control"
        id="floatingInput"
        v-model="email"
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
      <label>
        <input type="checkbox" value="remember-me" /> Remember me
      </label>
    </div>
    <button class="w-100 btn btn-lg btn-primary" type="button" v-on:click="doLogin">Sign in</button>
    <ul class="bottom-form-link">
      <li>
        <router-link to="/reset-password" class="link-secondary">Forgot Password?</router-link>
      </li>
      <li>
        <router-link to="/account/register" class="link-secondary">Sign up for an account</router-link>
      </li>
    </ul>

    <p class="mt-4 mb-3 text-muted">&copy; 2021</p>
  </form>
</template>

<script>
export default {
  name: "LoginPanel",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  props: {},
  created() {
    if (this.$route && this.$route.query && this.$route.query.userid) {

      this.email = this.$route.query.userid
    }
  },
  methods: {
    doLogin: function () {
      const axios = require("axios").default;

      axios
        .post(
          "http://localhost:8081/api/v1/account:signin",
          {
            username: this.email,
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

.login-divider {
  height: 0.1rem;
  background-color: rgba(0, 0, 0, 0.1);
}

.bottom-form-link {
    text-align: center;
    font-size: 14px;
    display: block;
    padding: 0;
}

.bottom-form-link li {
    display: inline-block;
    list-style: none;
}
.bottom-form-link ul {
    display: block;
    list-style-type: disc;
    margin-block-start: 1em;
    margin-block-end: 1em;
    margin-inline-start: 0px;
    margin-inline-end: 0px;
    padding-inline-start: 40px;
}

.bottom-form-link li:not(:first-child)::before {
  content: "\2022";
  margin: 0 8px 0px 4px;
}
</style>
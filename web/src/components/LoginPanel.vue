<template>
  <form class="form-signin my-3 px-3 align-items-center rounded-3 border shadow-lg">
    <img class="mb-4" src="../assets/WeekendClip.png" alt width="120" height="120" />
    <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

    <div class="form-floating">
      <input
        type="email"
        class="form-control"
        id="email"
        v-model="email"
        v-bind:class="[isValidEmail ? '' : invalidEmailClass]"
        placeholder="name@example.com"
        required
        @blur="checkEmail"
      />
      <label for="floatingInput">Email address</label>
      <div class="invalid-feedback">Email address is required</div>
    </div>
    <div class="form-floating py-1">
      <input
        type="password"
        class="form-control"
        id="password1"
        v-model="password"
        placeholder="Password"
        required
         @keyup.enter="doLogin"
      />
      <label for="floatingPassword">Password</label>
    </div>

    <div class="checkbox mb-3">
      <label>
        <input type="checkbox" value="remember-me" @keyup.enter="doLogin" /> Remember me
      </label>
    </div>
    <button class="w-100 btn btn-lg btn-primary" type="button" v-on:click="doLogin">Sign in</button>
    <div class="form-group">
      <div v-if="errors.length">
        <div
          v-for="error in errors"
          v-bind:key="error"
          class="alert alert-danger my-2"
          role="alert"
        >{{ error }}</div>
      </div>
    </div>
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
import { EventBus } from "@/event-bus";
export default {
  name: "LoginPanel",
  data() {
    return {
      validClass: "",
      password: "",
      isPasswordFilledIn: false,
      email: "",
      invalidEmailClass: "is-invalid",
      isValidEmail: true,
      errors: [],
    };
  },
  mounted() {
    var jwt = this.$cookies.get("authfulJwt");
    if (jwt != null) {
      this.$cookies.remove("authfulJwt");
      // TODO: Global event bus
      EventBus.emit("logout", "logout");
    }
  },
  created() {
    if (this.$route && this.$route.query && this.$route.query.userid) {
      this.email = this.$route.query.userid;
    }
  },
  methods: {
    checkEmail: function () {
      this.validClass = "is-valid";
      this.errors = [];
      if (this.email == "") {
        this.isValidEmail = false;
        this.invalidEmailClass = "";
      } else {
        if (this.validEmail(this.email)) {
          this.isValidEmail = true;
        } else {
          this.invalidEmailClass = "is-invalid";
          this.isValidEmail = false;
        }
      }
    },
    validEmail: function (email) {
      var re =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    },
    doLogin: function () {
      this.errors = [];

      this.checkEmail();
      if (!this.isValidEmail) {
        this.errors.push("Invalid e-mail address");
      }

      if (this.password.length == 0) {
        this.errors.push("Password is required");
      }

      if (this.errors.length) {
        return;
      }

      const axios = require("axios").default;

      axios
        .post(
          "http://localhost:8090/api/v1/account:signin",
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
            "authfulJwt",
            response.data.jwt,
            new Date(response.data.expires)
          );
          // Broadcast to the global event bus which will cause the navigation to evaluate the changes
          EventBus.emit("login", response.data.jwt);
          this.$router.push("/");
        })
        .catch((error) => {
          if (error.response.data) {
            if (error.response.data.error) {
              this.errors.push(error.response.data.error);

              return;
            }
          }
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
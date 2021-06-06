<template>
  <form
    class="form-register col-5 my-3 px-3 align-items-center rounded-3 border shadow-lg needs-validation"
    novalidate
    @submit="checkForm"
  >
    <img class="mb-4" src="../assets/WeekendClip.png" alt width="120" height="120" />
    <h1 class="h3 mb-3 fw-normal">Register for an Account</h1>
    <div class="form-group py-3">
      <div class="form-floating">
        <input
          type="email"
          class="form-control"
          id="email"
          v-model="email"
          v-bind:class="[isValidEmail ? validClass : invalidEmailClass]"
          placeholder="name@example.com"
          required
          @input="checkEmail"
        />
        <label for="floatingInput">Email address</label>
        <div class="invalid-feedback">Email address is required</div>
      </div>
    </div>
    <div class="form-group">
      <div class="form-floating py-1">
        <input
          type="password"
          class="form-control"
          v-bind:class="[isPasswordMatch ? validClass : invalidPasswordClass]"
          id="password1"
          v-model="password"
          placeholder="Password"
          required
          @input="checkPassword"
        />
        <label for="floatingPassword">Password</label>
      </div>
      <div class="form-floating py-1">
        <input
          type="password"
          class="form-control"
          v-bind:class="[isPasswordMatch ? validClass : invalidPasswordClass]"
          id="password2"
          v-model="confirmPassword"
          placeholder="Password"
          required
          @input="checkPassword"
        />
        <label for="floatingConfirmPassword">Re-enter password</label>
      </div>
    </div>
    <div class="form-group">
      <div v-if="errors.length">
        <b>There was a problem:</b>
        <div
          v-for="error in errors"
          v-bind:key="error"
          class="alert alert-warning"
          role="alert"
        >{{ error }}</div>
      </div>
    </div>
    <div class="form-group py-3">
      <button
        class="w-100 btn btn-lg btn-primary"
        v-bind:class="[isPasswordMatch && isValidEmail ? '' : 'disabled']"
        @click="doRegister"
        ref="createAccountButton"
        type="button"
      >Register</button>
      <div class="mt-2 mb-3">
        Already have an account?
        <router-link to="/login" class="link-secondary">Login</router-link>
      </div>
    </div>
    <p class="mt-4 mb-3 text-muted">&copy; 2021</p>
  </form>
</template>

<script>
export default {
  name: "LoginPanel",
  data() {
    return {
      validClass: "is-valid",
      password: "",
      confirmPassword: "",
      isPasswordMatch: false,
      invalidPasswordClass: "",
      email: "",
      invalidEmailClass: "",
      isValidEmail: false,
      errors: [],
    };
  },
  props: {},
  methods: {
    checkForm: function (e) {
      this.errors = [];

      if (!this.email) {
        this.errors.push("Email required.");
      } else if (!this.validEmail(this.email)) {
        this.errors.push("Valid email required.");
      }

      this.checkPassword();
      if (this.isPasswordMatch)
        if (!this.errors.length) {
          return true;
        }

      e.preventDefault();
      var form = document.querySelector(`form`);
      form.classList.add("was-validated");
    },
    checkEmail: function () {
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
    checkPassword: function () {
      if (
        this.password != this.confirmPassword ||
        this.password == "" ||
        this.confirmPassword == ""
      ) {
        this.isPasswordMatch = false;
        if (this.password == "") {
          this.invalidPasswordClass = "";
        } else {
          this.invalidPasswordClass = "is-invalid";
        }
      } else {
        this.isPasswordMatch = true;
      }
    },
    // global function

    doRegister: function () {
      this.checkForm();
      if (this.isPasswordMatch == false || this.isValidEmail == false) {
        return;
      }

      const axios = require("axios").default;

      axios
        .post(
          "http://localhost:8081/api/v1/account:signup",
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
          if (response.data.user_id) {
            if (response.data.user_id != "") {
              this.$router.push("/login?userid=" + this.email);
              return;
            }
          }

          if (response.data.error) {
            this.errors.push(response.data.error);
          } else {
            this.errors.push(
              "Unable to create user account. Please try again."
            );
          }
        })
        .catch((error) => {
          if (error.response.data) {
            if (error.response.data.error) {
              console.log(error.response.data);
              this.errors.push(error.response.data.error);
              return;
            }
          }
          console.log(error);
          this.errors.push("Unable to create user account. Please try again.");
        });
    },
  },
};
</script>

<style>
.form-register {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-register .checkbox {
  font-weight: 400;
}

.form-register .form-floating:focus-within {
  z-index: 2;
}

.form-register input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-register input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
<template>
  <form
    class="form-register col-5 my-3 px-3 align-items-center rounded-3 border shadow-lg needs-validation" novalidate  @submit="checkForm"
  >
    <img
      class="mb-4"
      src="../assets/WeekendClip.png"
      alt
      width="120"
      height="120"
    />
    <h1 class="h3 mb-3 fw-normal">Register for an Account</h1>
    <div class="form-group py-3">
      <div class="form-floating">
        <input
          type="email"
          class="form-control"
          id="floatingInput"
          v-model="email"
          placeholder="name@example.com"
          required
        />
        <label for="floatingInput">Email address</label>
         <div class="invalid-feedback">
      Email address is required
    </div>
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
          required
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
          required
        />
        <label for="floatingConfirmPassword">Re-enter password</label>
       
      </div>
    </div>
    <div class="form-group py-3">
      <button
        class="w-100 btn btn-lg btn-primary"
        type="submit"
      >
        Register
      </button>
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
      email: "",
      password: "",
      confirmPassword: "",
    };
  },
  props: {},
  methods: {
    // global function
    checkForm: function(e) {
      var form = document.querySelector(`form`);
  // Loop over them and prevent submission
      console.log(e);
      if (!form.checkValidity()) {
        console.log("not valid")
        e.preventDefault();
        e.target.checkValidity();      
      } else {
        console.log("valid")
      }
       form.classList.add('was-validated')
    //  if (!form.checkValidity()) {
    //       e.preventDefault()
    //       e.stopPropagation()
    //     } else {
    //         e.preventDefault()
    //         e.stopPropagation()       
    //     }
    },
    doRegister: function () {
      const axios = require("axios").default;

      axios
        .post(
          "http://localhost:8081/api/v1/account:resetPassword",
          {
            email: this.email,
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
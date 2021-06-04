<template>
  <div class="container-fluid">
    <div class="row justify-content-center">
      <form
        class="col-5 my-3 px-3 align-items-center rounded-3 border shadow-lg"
      >
        <h1 class="h3 my-4 mb-3 fw-normal">Trouble accessing your account?</h1>
        <p>Enter the email address you use for Hulu and we'll send you a password reset link.</p>
        <div class="form-group py-3 align-items-center">
          <div class="form-floating center">
            <input
              type="email"
              class="form-control"
              id="floatingInput"
              v-model="email"
              placeholder="name@example.com"
            />
            <label for="floatingInput">Email address</label>
          </div>
        </div>
         <div class="form-group py-3 ">
         <button
            class="w-50 btn btn-lg btn-primary "
            type="button"
            v-on:click="doReset"
          >
            Reset password
          </button>
          
          <ul class="bottom-form-link my-2">
            <li>
              <router-link to="/login" class="link-secondary"
                >Return to login</router-link
              >
            </li>
            <li>
              <router-link to="/account/register" class="link-secondary"
                >Sign up for an account</router-link
              >
            </li>
          </ul>
        </div>
        <p class="mt-4 mb-3 text-muted">&copy; 2021</p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "reset-password",
  props: {},
 data() {
    return {
      email: "",      
    };
 },
  methods: {
 doReset: function () {

  }

  },
};
</script>

<style>
</style>
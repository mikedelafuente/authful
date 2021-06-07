<template>
  <div>
    <router-link
      v-if="!isValidJwt"
      class="btn btn-bd-login d-lg-inline-block my-2 my-md-0 ms-md-3"
      active-class="active"
      to="/login"
    >Login</router-link>
    <router-link
      v-else
      class="btn btn-bd-login d-lg-inline-block my-2 my-md-0 ms-md-3"
      active-class="active"
      to="/logout"
    >Logout</router-link>  
  </div>
</template>

<script>
export default {
  data() {
    return {}
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
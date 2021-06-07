<template>
  <div>Logout</div>
</template>
<script>
import { EventBus } from "@/event-bus";
export default {
  name: "Logout",
  components: {},
  mounted() {
    this.$cookies.remove("authfulJwt");
    // TODO: Global event bus
    EventBus.emit("logout", "logout");
    this.$router.push("/login");
  },
  methods: {},
};
</script>

<style>
</style>
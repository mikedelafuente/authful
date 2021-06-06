<template>
  <div class="px-2 py-2 my-2 text-center">
    <div class="display-5 fw-bold">{{ msg }}</div>
    <div class="col-lg-6 mx-auto">
      <p class="lead mb-3">
        For a guide and recipes on how to configure / customize this project,
        <br />check out the
        <a
          href="https://cli.vuejs.org"
          target="_blank"
          rel="noopener"
        >vue-cli documentation</a>.
      </p>
    </div>
  </div>
  <div class="px-2 py-2 my-2 text-center">
    <div class="display-6 fw-bold">Installed CLI Plugins</div>
    <div class="col-lg-6 mx-auto">
      <ul>
        <li>
          <a
            href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel"
            target="_blank"
            rel="noopener"
          >babel</a>
        </li>
        <li>
          <a
            href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint"
            target="_blank"
            rel="noopener"
          >eslint</a>
        </li>
      </ul>
    </div>
  </div>
  <div class="px-2 py-1 my-1 text-center">
    <div class="display-6 fw-bold">Essential Links</div>
    <div class="col-lg-6 mx-auto">
      <ul>
        <li>
          <a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a>
        </li>
        <li>
          <a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a>
        </li>
        <li>
          <a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a>
        </li>
        <li>
          <a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a>
        </li>
        <li>
          <a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a>
        </li>
      </ul>
    </div>
  </div>

  <div class="px-2 py-2 my-2 text-center">
    <div class="display-6 fw-bold">Ecosystem</div>
    <div class="col-lg-6 mx-auto">
      <ul>
        <li>
          <a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a>
        </li>
        <li>
          <a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a>
        </li>
        <li>
          <a
            href="https://github.com/vuejs/vue-devtools#vue-devtools"
            target="_blank"
            rel="noopener"
          >vue-devtools</a>
        </li>
        <li>
          <a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a>
        </li>
        <li>
          <a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a>
        </li>
      </ul>
    </div>
  </div>

  <button class="btn btn-primary btn-lg px-4 gap-3" v-on:click="reverseMessage">Reverse Message</button>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
  methods: {
    reverseMessage: function () {
      this.$.props.msg = this.$.props.msg.split("").reverse().join("");     
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

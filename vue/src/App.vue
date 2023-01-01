<template>
  <div style="margin-bottom:10px">
    <SpTheme>
      <SpNavbar class="navbar" :links="navbarLinks" :active-route="router.currentRoute.value.path" />
      <router-view />
    </SpTheme>
  </div>
</template>

<script lang="ts">
import { SpNavbar, SpTheme } from "@starport/vue";
import { computed, onBeforeMount } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";

export default {
  components: { SpTheme, SpNavbar },

  setup() {
    // store
    let $s = useStore();

    // router
    let router = useRouter();

    // state
    let navbarLinks = [
      { name: "Portfolio", url: "/portfolio" },
      { name: "Employers", url: "/employers" },
    ];

    // computed
    let address = computed(() => $s.getters["common/wallet/address"]);

    // lh
    onBeforeMount(async () => {
      await $s.dispatch("common/env/init");

      // router.push("portfolio");
    });

    return {
      navbarLinks,
      // router
      router,
      // computed
      address,
    };
  },
};
</script>

<style scoped lang="scss">
body {
  margin: 0;
}
</style>

<style lang="scss">
#app {
  .navbar-wrapper[data-v-6b5cb40d] {
    display: flex;
    justify-content: space-between;
    height: 80px;
    left: 0;
    right: 0;
    top: 0;
    background: #ffffff;
    margin-bottom: 0px;
  }
}
</style>


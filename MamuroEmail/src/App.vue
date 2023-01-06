<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import CardComponent from "./components/CardComponent.vue";
import SearchComponent from "./components/SearchComponent.vue";
import PaginationComponent from "./components/PaginationComponent.vue";
</script>

<template>
  <div class="contenedor flex items-center justify-center flex-col gap-5">
    <h1 className="text-3xl">Mamuro Email</h1>
    <SearchComponent />
  </div>
  <!-- <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView /> -->
</template>

<style scoped></style>

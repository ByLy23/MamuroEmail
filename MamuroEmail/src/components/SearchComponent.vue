<template>
  <link
    rel="stylesheet"
    href="https://unpkg.com/flowbite@1.4.4/dist/flowbite.min.css"
  />

  <div class="max-w-full mx-auto w-1/2">
    <form class="flex items-center">
      <label for="simple-search" class="sr-only">Search</label>
      <div class="relative w-full">
        <div
          class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none"
        ></div>
        <input
          type="text"
          v-model="searchQuery"
          id="simple-search"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Search"
          required
        />
      </div>
      <button
        type="button"
        @click="search(1)"
        class="p-2.5 ml-2 text-sm font-medium text-white bg-[#d96c06] rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          ></path>
        </svg>
      </button>
    </form>
  </div>
  <div class="mt-10 flex items-center justify-center flex-col gap-5">
    <CardComponent
      :items="items"
      :increment="increment"
      :maxPage="maxPage"
      :currentPage="currentPage"
      :decrement="decrement"
    />
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { getSearch } from "../api/getSearch";
import CardComponent from "./CardComponent.vue";
import PaginationComponent from "./PaginationComponent.vue";
export default defineComponent({
  name: "SearchComponent",
  data() {
    return {
      searchQuery: "",
      items: [],
      increment: this.incrementPage,
      currentPage: 1,
      maxPage: 1,
      decrement: this.decrementPage,
    };
  },
  methods: {
    async search(currentPage = 1) {
      const search = await getSearch(this.searchQuery, currentPage);
      const { total, hits } = search.hits;
      localStorage.setItem(
        "search_elements",
        JSON.stringify({
          maxPages: Math.ceil(total.value / 20),
          currentPage: currentPage,
        })
      );
      console.log(hits);
      this.items = hits;
      this.maxPage = Math.ceil(total.value / 20);
      // this.items = search;
    },
    incrementPage() {
      const infoPages = localStorage.getItem("search_elements") || "";
      const { maxPages, currentPage } = JSON.parse(infoPages);
      if (currentPage < maxPages) {
        localStorage.setItem(
          "search_elements",
          JSON.stringify({
            maxPages,
            currentPage: currentPage + 1,
          })
        );
        this.currentPage = currentPage + 1;

        this.search(currentPage + 1);
      }
    },
    decrementPage() {
      const infoPages = localStorage.getItem("search_elements") || "";
      const { maxPages, currentPage } = JSON.parse(infoPages);
      console.log("a");
      if (currentPage > 1) {
        localStorage.setItem(
          "search_elements",
          JSON.stringify({
            maxPages,
            currentPage: currentPage - 1,
          })
        );
        this.currentPage = currentPage - 1;
        this.search(currentPage - 1);
      }
    },
  },
  components: { CardComponent, PaginationComponent },
});
</script>

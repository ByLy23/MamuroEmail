<template>
  <!-- This is an example component -->
  <div class="flex items-center justify-center mb-10">
    <div class="flex justify-center items-center space-x-4">
      <button
        @click="decrement"
        class="border rounded-md bg-gray-100 px-2 py-1 text-3xl leading-6 text-slate-400 transition hover:bg-gray-200 hover:text-slate-500 cursor-pointer shadow-sm"
      >
        &lt;
      </button>
      <div class="text-slate-500">{{ currentPage }} /{{ maxPage }}</div>
      <button
        @click="increment"
        class="border rounded-md bg-gray-100 px-2 py-1 text-3xl leading-6 text-slate-400 transition hover:bg-gray-200 hover:text-slate-500 cursor-pointer shadow-sm"
      >
        >
      </button>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props: ["maxPage", "currentPage", "items", "increment", "decrement"],
  data() {
    return {
      currentPage: this.currentPage,
      maxPage: this.maxPage,
      increment: this.increment,
      decrement: this.decrement,
    };
  },
};
</script>

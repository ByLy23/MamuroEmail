<script setup lang="ts">
import CardItemComponent from "./CardItemComponent.vue";
const items = [
  {
    title: "Allen-p",
    content: "Message from to the future...",
    file: "allen-p.png",
    root: "allen-p",
    date: "08/10/2022",
  },
  {
    title: "Al2len-p",
    content: "Message from to the future...",
    file: "alle2n-p.png",
    root: "al2len-p",
    date: "082/10/2022",
  },
];
</script>
<template>
  <ul class="flex flex-row gap-3 flex-wrap ml-7 mr-7 justify-center mb-10">
    <li v-for="item in items" :key="item.title">
      <CardItemComponent
        :title="item.title"
        :content="item.content"
        :root="item.root"
        :file="item.file"
        :date="item.date"
      />
    </li>
  </ul>
</template>

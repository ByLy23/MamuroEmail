<script setup lang="ts">
import { defineComponent } from "vue";
import CardItemComponent from "./CardItemComponent.vue";
import PaginationComponent from "./PaginationComponent.vue";
</script>
<template>
  <ul class="flex flex-row gap-3 flex-wrap ml-7 mr-7 justify-center mb-10">
    <li v-for="item in items" :key="item._id">
      <CardItemComponent
        :title="item._source.username"
        :content="item._source.content"
        :root="item._source.directory"
        :file="item._source.file_name"
      />
    </li>
  </ul>
  <div v-if="items.length !== 0">
    <PaginationComponent
      :items="items"
      :increment="increment"
      :currentPage="currentPage"
      :maxPage="maxPage"
      :decrement="decrement"
    />
  </div>
</template>
<script lang="ts">
export default defineComponent({
  name: "CardComponent",
  components: {
    CardItemComponent,
    PaginationComponent,
  },
  props: ["items", "increment", "currentPage", "maxPage", "decrement"],
  data() {
    return {
      items: this.items,
      increment: this.increment,
      decrement: this.decrement,
      currentPage: this.currentPage,
      maxPage: this.maxPage,
    };
  },
});
</script>

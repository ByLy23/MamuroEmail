<script setup lang="ts">
import { defineProps } from "vue";
import ModalComponent from "./ModalComponent.vue";
defineProps<{
  title: string;
  content: string;
  file: string;
  root: string;
  date: string;
}>();
</script>
<template>
  <article class="bg-white shadow-2xl rounded-2xl p-5 w-auto max-w-sm">
    <h1 class="font-bold text-yellow-500">{{ title }}</h1>
    <h3 class="font-bold text-yellow-200">{{ file }}</h3>
    <p class="font-light text-gray-500 hover:font-bold">
      {{ content.substring(0, 100) }}
    </p>
    <h6 class="text-sm text-gray-300 mb-5">Encontrado: {{ date }}</h6>
    <button
      class="block text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      type="button"
      :data-modal-toggle="'default-modal' + file"
    >
      Ver Mas...
    </button>
    <ModalComponent
      :title="title"
      :content="content"
      :file="'default-modal' + file"
    />
  </article>
</template>

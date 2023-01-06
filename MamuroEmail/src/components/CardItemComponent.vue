<script setup lang="ts">
import { defineProps } from "vue";
defineProps<{
  title: string;
  content: string;
  file: string;
  root: string;
}>();
</script>
<template>
  <!-- component -->
  <!-- This is an example component -->
  <div class="flex flex-col border overflow-hidden bg-white">
    <div class="grid grid-cols-1 sm:grid-cols-4">
      <div
        class="flex flex-col border-b sm:border-b-0 items-center p-8 sm:px-4 sm:h-full sm:justify-center"
      >
        <p class="text-7xl font-bold text-[#0ed3cf] rounded-full">
          {{ title.substring(0, 1).toUpperCase() }}
        </p>
      </div>
      <div class="flex flex-col sm:border-l col-span-3">
        <div class="flex flex-col space-y-4 p-6 text-gray-600">
          <div class="flex flex-row text-sm">
            <span class="mr-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z"
                />
              </svg>
            </span>
            <p class="flex items-center text-gray-500">
              <span class="font-semibold mr-2 text-xs uppercase">Name:</span>
              <span>{{ title }}</span>
            </p>
          </div>

          <div class="flex flex-row text-sm">
            <span class="mr-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.25 12.75V12A2.25 2.25 0 014.5 9.75h15A2.25 2.25 0 0121.75 12v.75m-8.69-6.44l-2.12-2.12a1.5 1.5 0 00-1.061-.44H4.5A2.25 2.25 0 002.25 6v12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9a2.25 2.25 0 00-2.25-2.25h-5.379a1.5 1.5 0 01-1.06-.44z"
                />
              </svg>
            </span>
            <p class="flex items-center text-gray-500">
              <span class="font-semibold mr-2 text-xs uppercase"
                >File root:</span
              >
              <span>{{ root }}</span>
            </p>
          </div>
          <div class="flex flex-row text-sm">
            <span class="mr-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z"
                />
              </svg>
            </span>
            <p class="flex items-center text-gray-500">
              <span class="font-semibold mr-2 text-xs uppercase"
                >File Name:</span
              >
              <span>{{ file }}</span>
            </p>
          </div>
        </div>
        <div class="flex flex-col w-full relative bottom-0"></div>
      </div>
    </div>
  </div>

  <div class="flex flex-col overflow-hidden bg-white my-0 max-w-lg">
    <div class="grid grid-cols-1">
      <div class="flex flex-col">
        <div class="flex flex-col space-y-4 p-6 text-gray-600">
          <div class="flex flex-row font_small">
            <p class="flex items-center p-1 text-gray-500">
              {{ content }}
            </p>
          </div>
        </div>
        <div class="flex flex-col w-full relative bottom-0"></div>
      </div>
    </div>
  </div>
</template>

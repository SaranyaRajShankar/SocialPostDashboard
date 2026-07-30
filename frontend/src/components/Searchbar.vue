<template>
  <div>
    <input v-model="searchQuery" type="text" placeholder="search..." />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from "vue";

const searchQuery = ref("");
const emit = defineEmits(["updateSearchQuery"]);


watch(searchQuery, (newValue) => {
  emit("updateSearchQuery", newValue);
});
</script>

<style></style>

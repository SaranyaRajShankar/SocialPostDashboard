<template>
  <div>
    <searchbar @updateSearchQuery="updateSearchQuery" />
    <Table :searchQuery="searchQuery" />
  </div>
</template>

<script setup lang="ts">
import Table from "./components/Table.vue";
import Searchbar from "./components/Searchbar.vue";
import { ref, watch } from "vue";

const searchQuery = ref("");
const updateSearchQuery = (newValue: string) => {
  searchQuery.value = newValue;
};
</script>

<style></style>

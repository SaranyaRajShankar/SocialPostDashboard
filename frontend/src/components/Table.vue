<template>
  <div>
    <table class="table-container">
      <thead>
        <tr>
          <th>UserId</th>
          <th>Id</th>
          <th
            @click="sortPosts"
            @mouseenter="hover = true"
            @mouseleave="hover = false"
          >
            Title {{ hover ? "^" : "" }}
          </th>
          <th>Body</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="post in props.searchQuery ? filteredPosts : postsPerPage"
          :key="post.id"
        >
          <td>{{ post.userId }}</td>
          <td>{{ post.id }}</td>
          <td class="post-content">{{ post.title }}</td>
          <td class="post-content">{{ post.body }}</td>
        </tr>
      </tbody>
    </table>
    <!-- <div class="button-group">
      <button @click="decrementPage()" :disabled="startIndex === 0">
        Prev
      </button>
      <button
        @click="incrementPage()"
        :disabled="startIndex + pageSize === posts.length"
      >
        Next
      </button>
    </div> -->
    <div ref="loadMore"></div>
  </div>
</template>
<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";

const startIndex = ref(0);
const pageSize = ref(5);
const visiblePosts = ref(5);
const hover = ref(false);
const postsPerPage = computed(() => {
  return posts.value.slice(0, visiblePosts.value);
});

const props = defineProps({
  searchQuery: {
    type: String,
    required: true,
  },
});

const filteredPosts = ref<Post[]>([]);

const sortPosts = () => {
  if (props.searchQuery) {
    filteredPosts.value.sort((a, b) => a.title.localeCompare(b.title));
  } else {
    posts.value.sort((a, b) => a.title.localeCompare(b.title));
  }
};

watch(
  () => props.searchQuery,
  (newValue) => {
    if (newValue) {
      filteredPosts.value = posts.value.filter((post) =>
        post.title.toLowerCase().includes(newValue.toLowerCase()),
      );
      visiblePosts.value = filteredPosts.value.length;
    } else {
      visiblePosts.value = 5;
    }
    console.log(filteredPosts.value.length);
  },
);

interface Post {
  userId: number;
  id: number;
  title: string;
  body: string;
}

const posts = ref<Post[]>([]);

const loadMore = ref<HTMLElement | null>(null);

let observer: IntersectionObserver;

onMounted(async () => {
  const response = await fetch("https://jsonplaceholder.typicode.com/posts");
  posts.value = await response.json();

  observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      loadNextPosts();
    }
  });
  if (loadMore.value) {
    observer.observe(loadMore.value);
  }
});

const loadNextPosts = () => {
  if (visiblePosts.value < posts.value.length) {
    visiblePosts.value += 5;
  }
};

const decrementPage = () => {
  if (startIndex.value > 0) startIndex.value = startIndex.value - 5;
};

const incrementPage = () => {
  if (startIndex.value + pageSize.value < posts.value.length)
    startIndex.value = startIndex.value + 5;
};
</script>
<style>
.table-container {
  width: 100%;
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px auto;
  font-family: Arial sans-serif;
  font-size: 14px;
}
thead th {
  padding: 20px;
  background-color: rgb(23, 105, 172);
}
tbody td {
  padding: 20px;
  background-color: rgb(1, 23, 41);
}
.post-content {
  text-align: left;
  max-width: 200px;
  word-wrap: break-word;
}
.button-group {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>

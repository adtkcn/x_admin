<template>
  <uv-load-more
    :status="status"
    :loading-text="loadingText"
    :loadmore-text="loadmoreText"
    :nomore-text="nomoreText"
    @loadmore="loadmore"
  />
</template>
<script setup lang="ts">
// defineEmits 为编译宏，无需 import
defineProps({
  status: {
    type: String,
    default: "loadmore",
  },
  loadingText: {
    type: String,
    default: "努力加载中",
  },
  loadmoreText: {
    type: String,
    default: "轻轻上拉",
  },
  nomoreText: {
    type: String,
    default: "实在没有了",
  },
});

const emits = defineEmits(["loadmore"]);

function loadmore(e: any) {
  emits("loadmore", e);
}
</script>

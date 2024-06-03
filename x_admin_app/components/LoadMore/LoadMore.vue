<template>
  <uv-load-more
    :status="props.status"
    :loading-text="props.loadingText"
    :loadmore-text="props.loadmoreText"
    :nomore-text="props.nomoreText"
    @loadmore="loadmore"
  />
</template>
<script setup>
import { onReady, onPullDownRefresh, onReachBottom } from "@dcloudio/uni-app";

import { ref,  defineEmits,defineModel } from "vue";

var pager = defineModel();

const emits = defineEmits(["loadmore"]);
const props = defineProps({
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
function loadmore(e) {
  emits("loadmore", e);
}
</script>

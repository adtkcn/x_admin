<template>
<view>

  <uv-sticky :customNavHeight="0" bgColor="#fff">
    <uv-status-bar></uv-status-bar>
    <uv-navbar
      leftText=""
      :safeAreaInsetTop="false"
      :fixed="false"
      title="{{{.FunctionName}}}"
      autoBack
    >
      <template v-slot:right>
        <uv-icon v-if="!fromSearch" name="search" size="24" @click="moreSearch"></uv-icon>
      </template>
    </uv-navbar>
    <view class="search">
      <uv-search
        placeholder="请输入搜索内容"
        shape="square"
        v-model="queryParams.key"
        :showAction="false"
        bgColor="#fff"
        borderColor="rgba(0, 0, 0, .1)"
        @search="resetPage"
      ></uv-search>
    </view>
  </uv-sticky>

  <uv-list>
    <uv-list-item
      v-for="item of pager.lists"
      :key="item.id"
      clickable
      show-arrow
      :title="item.name"
      :note="item.placedPostion+'('+item.placedName+')'"
      :right-text="item.status == 1 ? '在线' : ''"
      @click="toDetails(item)"
    ></uv-list-item>
  </uv-list>

  <uv-back-top :scroll-top="scrollTop"></uv-back-top>

  <uv-loading-page
    :loading="pager.pageNo == 1 && pager.loading == 'loading'"
    loading-text="加载中..."
    font-size="24rpx"
  ></uv-loading-page>
  <uv-load-more
    :status="pager.loading"
    :loading-text="pager.loadingText"
    :loadmore-text="pager.loadmoreText"
    :nomore-text="pager.nomoreText"
    @loadmore="NextPage"
  />
</view>
</template>

<script setup>
import { reactive, ref } from "vue";
import {
  onLoad,
  onPullDownRefresh,
  onReachBottom,
  onPageScroll,
} from "@dcloudio/uni-app";
import { equipment_list } from "@/api/equipment.js";
import { usePaging } from "@/utils/usePaging";
import { toPath } from "@/utils/utils";
const queryParams = reactive({
  key: "",
  
});
let fromSearch=ref(false);
onLoad((e) => {
  console.log("equipment onLoad", e);
  if (e) {
    for (const key in e) {
      if (Object.hasOwnProperty.call(e, key)) {
        fromSearch.value = true;
        queryParams[key] = e[key];
      }
    }
  }
  getLists();
});
const { pager, getLists, NextPage, resetPage, resetParams } = usePaging({
  fetchFun: equipment_list,
  params: queryParams,
});
let scrollTop = ref(0);
onPageScroll((e) => {
  scrollTop.value = e.scrollTop;
});
onPullDownRefresh(() => {
  resetPage();
});
onReachBottom(() => {
  NextPage();
});

function toDetails(item) {
  toPath("/pages/equipment/details", { id: item.id });
}
function moreSearch() {
  toPath("/pages/equipment/search");
}
</script>

<style lang="scss" scoped>
.search {
  padding: 5rpx;
  background-color: #fff;
}
</style>

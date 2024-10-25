<template>
<view>
<!--
  <uv-sticky :customNavHeight="0" bgColor="#fff">
    <uv-status-bar></uv-status-bar>
    <uv-navbar
      leftText=""
      :safeAreaInsetTop="false"
      :fixed="false"
      title="用户协议"
      autoBack
    >
      <template v-slot:right>
        <uv-icon v-if="!fromSearch" name="search" size="24" @click="moreSearch"></uv-icon>
      </template>
    </uv-navbar>
 
  </uv-sticky>
 -->
  <uv-list>
    <uv-list-item
      v-for="item of pager.lists"
      :key="item.Id"
      clickable
      show-arrow
      :title="item.Id"
      :right-text="item.Id"
      @click="toDetails(item)"
    ></uv-list-item>
  </uv-list>
  <wd-fab v-model:active="activeFab" :draggable="true">
      <wd-button v-if="!fromSearch" custom-class="fab-button" type="primary" round @click="moreSearch" >
        <wd-icon name="search" size="20px"></wd-icon>
      </wd-button>
      <wd-button v-if="$perms('admin:user_protocol:add')"  custom-class="fab-button" type="primary" round @click="add">
        <wd-icon name="add" size="20px"></wd-icon>
      </wd-button>
  </wd-fab>
  <uv-back-top :scroll-top="scrollTop"></uv-back-top>

  <uv-empty v-if="pager.loading =='nomore'&&pager.lists.length == 0" marginTop="150" mode="data"></uv-empty>
  <uv-loading-page
    :loading="pager.pageNo == 1 && pager.loading == 'loading'"
    loading-text="加载中..."
    font-size="24rpx"
  ></uv-loading-page>
  <uv-load-more
    v-if="pager.lists.length > 0"
    :status="pager.loading"
    :loading-text="pager.loadingText"
    :loadmore-text="pager.loadmoreText"
    :nomore-text="pager.nomoreText"
    @loadmore="NextPage"
  />
</view>
</template>

<script setup  lang="ts">
import { reactive, ref } from "vue";
import {
  onLoad,
  onPullDownRefresh,
  onReachBottom,
  onPageScroll,
} from "@dcloudio/uni-app";
import { user_protocol_list } from "@/api/user/protocol";
import type { type_user_protocol,type_user_protocol_query	} from "@/api/user/protocol";

import { usePaging } from "@/hooks/usePaging";
import { toPath } from "@/utils/utils";
const queryParams = reactive<type_user_protocol_query>({
    Title: '',
    Content: '',
    Sort: '',
    CreateTimeStart: '',
    CreateTimeEnd: '',
    UpdateTimeStart: '',
    UpdateTimeEnd: '',
});
let activeFab = ref(false);
let fromSearch=ref(false);
onLoad((e) => {
  console.log("user_protocol onLoad", e);
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
const { pager, getLists, NextPage, resetPage, resetParams } = usePaging<type_user_protocol>({
  fetchFun: user_protocol_list,
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
  toPath("/pages/user/protocol/details", { Id: item.Id });
}
function moreSearch() {
  toPath("/pages/user/protocol/search");
}
function add() {
  toPath("/pages/user/protocol/edit");
}
</script>

<style lang="scss" scoped>
  :deep(.fab-button) {
    min-width: auto !important;
    box-sizing: border-box;
    width: 40px !important;
    height: 40px !important;
    border-radius: 40px !important;
    margin: 8rpx;
  }
</style>

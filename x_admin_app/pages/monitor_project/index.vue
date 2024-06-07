<template>
  <view>
  
    <uv-sticky :customNavHeight="0" bgColor="#fff">
      <uv-status-bar></uv-status-bar>
      <uv-navbar
        leftText=""
        :safeAreaInsetTop="false"
        :fixed="false"
        title="监控项目"
        autoBack
      >
        <template v-slot:right>
          <uv-icon v-if="!fromSearch" name="search" size="24" @click="moreSearch"></uv-icon>
        </template>
      </uv-navbar>
   
    </uv-sticky>
  
    <uv-list>
      <uv-list-item
        v-for="item of pager.lists"
        :key="item.id"
        clickable
        show-arrow
        :title="item.id"
        :note="item.id"
        :right-text="item.id"
        @click="toDetails(item)"
      ></uv-list-item>
    </uv-list>
  
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
  import { monitor_project_list } from "@/api/monitor_project";
  import type { type_monitor_project,type_monitor_project_query	} from "@/api/monitor_project";
  
  import { usePaging } from "@/hooks/usePaging";
  import { toPath } from "@/utils/utils";
  const queryParams = reactive<type_monitor_project_query>({
      projectKey: '',
      projectName: '',
      projectType: '',
      createTimeStart: '',
      createTimeEnd: '',
      updateTimeStart: '',
      updateTimeEnd: '',
  });
  let fromSearch=ref(false);
  onLoad((e) => {
    console.log("monitor_project onLoad", e);
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
  const { pager, getLists, NextPage, resetPage, resetParams } = usePaging<type_monitor_project>({
    fetchFun: monitor_project_list,
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
    toPath("/pages/monitor_project/details", { id: item.id });
  }
  function moreSearch() {
    toPath("/pages/monitor_project/search");
  }
  </script>
  
  <style lang="scss" scoped>
   
  </style>
  
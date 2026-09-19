<template>
  <view class="x-card" :style="{ marginBottom: bottom + 'px' }">
    <view class="x-card-image" v-if="icon">
      <uv-image width="100%" :src="icon" mode="widthFix"></uv-image>
    </view>
    <view class="x-card-title" v-if="title">
      <text>{{ title }}</text>
    </view>
    <view class="x-card-content" v-if="content">     
      <slot name="contentSlot">
        <text>{{ content }}</text>
      </slot>
    </view>
    <view class="x-card-bottom" v-if="btnText">
      <slot name="btnSlot">
        <button :open-type="openType" @click="btnClick">{{ btnText }}</button>
      </slot>
    </view>
  </view>
</template>
<script setup lang="ts">
import { toPath } from "@/utils/utils";

const props = defineProps({
  icon: {
    default: "",
    type: [String, Number],
  },
  title: {
    default: "",
    type: [String, Number],
  },
  content: {
    default: "",
    type: [String, Number],
  },
  btnText: {
    default: "",
    type: [String, Number],
  },
  openType: {
    default: "",
    type: String,
  },
  bottom: {
    default: "0",
    type: [String, Number],
  },
  url: {
    default: "",
    type: String,
  },
});

const emit = defineEmits(["btnClick"]);

function btnClick() {
  if (props.url) {
    toPath(props.url as string);
    return;
  }
  emit("btnClick");
}
</script>
<style lang="scss">
.x-card {
  background-color: white;
  margin:0 10rpx;
  padding: 10px 10px;
  border-radius: 8px;
  // margin-bottom: 10px;

  .x-card-image {
    padding-bottom:  10px;
  }
  .x-card-title {
    padding-bottom:  10rpx;

    text {
      line-height: 1;
      font-size: 32rpx; 
    //   color: #aaaaaa;
    }
  }

  .x-card-content {
    padding-bottom:  10rpx;

    text {
      line-height: 1;
      font-size: 26rpx;
    //   color: #aaaaaa;
    }
  }

  .x-card-bottom {
    padding-top:  20rpx;
    button {
      margin: 0;
      padding: 16rpx 16rpx;
      line-height: 1;
      border: 0;
      font-size: 24rpx;
      background-color: rgba(111, 148, 205, 1);
      color: white;
    }
  }
}
</style>

<template>
  <view class="my">
    <uv-sticky offsetTop="-1" bgColor="#ffffff">
      <view style="background: #fff">
        <uv-cell
          :title="userStore?.username"
          :label="userStore?.nickname"
          isLink
        >
          <!--       -->
          <template v-slot:icon>
            <button class="avatar-wrapper" open-type="chooseAvatar" @click.stop>
              <uv-avatar
                height="80rpx"
                width="80rpx"
                shape="square"
                :src="$filePath(userStore.avatar)"
                mode="aspectFill"
              ></uv-avatar>
            </button>
          </template>
        </uv-cell>
      </view>
    </uv-sticky>
    <view class="btn">
      <uv-button type="primary" text="登录" @click="toLogin"></uv-button>
    </view>
  </view>
</template>

<script setup>
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();

function toLogin() {
  uni.redirectTo({
    url: "/pages/login/login",
  });
}
</script>

<style scoped>
.my {
  background-color: rgba(243, 246, 249, 1);
  padding-bottom: 40px;
  min-height: 100vh;
  box-sizing: border-box;
}

.avatar-wrapper {
  width: 80rpx;
  height: 80rpx;
  padding: 0px;
  background-color: white;
  border: 0;
  outline: none;
}

.btn {
  margin: 200rpx 20rpx;
}
</style>

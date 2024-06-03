<template>
  <view class="my">
    <uv-sticky offsetTop="-1" bgColor="#ffffff">
      <view style="background: #fff">
        <uv-cell :title="userStore?.username"  :label="userStore?.nickname" isLink>
          <!--       -->
          <template v-slot:icon>
            <button class="avatar-wrapper" open-type="chooseAvatar" @click.stop>
              <uv-avatar
                width="100%"
                height="100%"
                shape="circle"
                :src="$filePath(userStore.avatar)"
                mode="widthFix"
              ></uv-avatar>
            </button>
          </template>
        </uv-cell>
      </view>
    </uv-sticky>
    <uv-button
      class="btn"
      type="primary"
      text="登录"
      @click="toLogin"
    ></uv-button>
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
  width: 100%;
  height: 100%;
  padding: 7px;
  background-color: white;
  border: 0;
  outline: none;
}

.btn {
  margin: 200rpx 20rpx;
}
</style>

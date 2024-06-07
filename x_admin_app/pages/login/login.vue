<template>
  <view class="form">
    <uv-form ref="form" labelAlign="right" :model="model" :rules="rules">
      <uv-form-item :customStyle="hide">
        <uv-input
          v-model="model.username"
          placeholder="请输入账号"
          :rules="[{ required: true, message: '请填写账号' }]"
        />
      </uv-form-item>
      <uv-form-item :customStyle="hide">
        <uv-input
          type="password"
          v-model="model.password"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />
      </uv-form-item>
      <uv-form-item label="账号" name="username" prop="username">
        <uv-input
          v-model="model.username"
          placeholder="请输入账号"
          :rules="[{ required: true, message: '请填写账号' }]"
        />
      </uv-form-item>
      <uv-form-item label="密码" name="password" prop="password">
        <uv-input
          type="password"
          v-model="model.password"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />
      </uv-form-item>

      <Verify
        :mode="'pop'"
        :captchaType="'blockPuzzle'"
        ref="verifyRef"
        :imgSize="{ width: 310, height: 155 }"
        @success="VerifySuccess"
      ></Verify>

      <view class="footer">
        <uv-button type="primary" text="登录" @click="VerifyShow"></uv-button>
      </view>
    </uv-form>
  </view>
</template>

<script lang="ts" setup>
import { ref, reactive ,shallowRef} from "vue";

import uvForm from "@/uni_modules/uv-form/components/uv-form/uv-form.vue";
import Verify from "@/components/verify/verify.vue";

import { login } from "@/api/user";
import { setLocalStorage } from "@/utils/storage";
import { alert ,encryptPassword} from "@/utils/utils";
// import appicon from "@/static/appicon.png";

import { useUserStore } from "@/stores/user";
const hide = {
  overflow: "hidden",
  // visibility: 'hidden',
  height: 0,
  padding: 0,
  border: 0,
};
const userStore = useUserStore();

const model = reactive({
  username: "",
  password: "",
});
const rules = {
  username: [
    {
      required: true,
      message: "请填写账号",
      trigger: ["blur", "change"],
    },
  ],
  password: [
    {
      required: true,
      message: "请填写密码",
      trigger: ["blur", "change"],
    },
  ],
};
const form = ref();
// const verifyRef = ref();
const verifyRef = shallowRef<InstanceType<typeof Verify>>()

function VerifyShow() {
  form.value
    .validate()
    .then(() => {
      verifyRef.value.show();
    })
    .catch((err) => {
      // alert(err.message);
    });
}
function VerifySuccess(e) {
  console.log("VerifySuccess", e);
  handleSubmit(e);
}
function handleSubmit(verify) {
  let data = {
    username: model.username,
    password:(model.password),

    ...verify,
  };

  // userStore.getUser();

  login(data)
    .then((res) => {
      if (res.code === 200) {
        setLocalStorage("token", res.data.token);

        userStore.getInfo().then((res) => {
          uni.switchTab({
            url: "/pages/index/index",
          });
        });
      } else {
        alert(res.message);
      }
    })
    .catch((err) => {
      alert(err.message);
    });
}
</script>

<style lang="scss" scoped>
.form {
  margin: 200rpx 20rpx;

  .footer {
    padding: 60rpx 20rpx 200rpx;
  }
}
</style>

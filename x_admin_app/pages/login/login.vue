<template>
  <view>
    <uv-form
      class="form"
      ref="form"
      labelWidth="60"
      labelAlign="right"
      :model="model"
      :rules="rules"
    >
      <uv-form-item class="hide">
        <uv-input
          class="hide"
          v-model="model.username"
          placeholder="请输入账号"
          :rules="[{ required: true, message: '请填写账号' }]"
        />
      </uv-form-item>
      <uv-form-item class="hide">
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
          @success="VerifySuccess"
          :mode="'pop'"
          :captchaType="'blockPuzzle'"
          ref="verifyRef"
          :imgSize="{ width: '310px', height: '155px' }"
        ></Verify>
 

      <view class="footer">
        <!--  -->
        
        <wd-button class="btn" size="large" @click="VerifyShow" block
          >VerifyShow</wd-button
        >

        <wd-button class="btn" size="large" @click="handleSubmit" block
          >登录</wd-button
        >
      </view>
    </uv-form>
  </view>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { v1 as uuid_v1 } from "uuid";

import Verify from "@/components/verifition/verify";

import { login } from "@/api/user.js";
import { setLocalStorage } from "@/utils/storage.js";
import { toast } from "@/utils/utils.js";
// import appicon from "@/static/appicon.png";

import { useUserStore } from "@/stores/user";

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
const verifyRef=ref()
function VerifyShow(e) {
  verifyRef.value.show()
}
function VerifySuccess(e) {
  console.log('VerifySuccess',e);
  
  handleSubmit(e)
}
function handleSubmit(verify) {
  form.value
    .validate()
    .then((res) => {
      let data = {
        username: model.username,
        password: model.password,
 
        ...verify
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
  
            toast(res.message);
          }
        })
        .catch((err) => {
          toast(err.message);
        });
    })
    .catch((err) => {
      toast(err.message);
    });
}
</script>

<style lang="scss" scoped>
.hide {
  height: 0 !important;
  padding: 0 !important;

  border: 0 !important;
  overflow: hidden !important;
}

.form {
  margin: 200rpx 20rpx;

  .captchaImage {
    padding: 0 10rpx;
  }

  .footer {
    padding: 60rpx 20rpx 200rpx;

    .btn {
      background-color: #00b294;
    }
  }
}
</style>

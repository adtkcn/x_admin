<template>
  <view class="page-content">
    <uv-form
      labelPosition="left"
      :model="form"
      :rules="formRules"
      ref="formRef"
    >
      <uv-form-item label="项目key" prop="project_key" borderBottom>
        <uv-input v-model="form.project_key" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="sdk生成的客户端id" prop="client_id" borderBottom>
        <uv-input v-model="form.client_id" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="用户id" prop="user_id" borderBottom>
        <uv-input v-model="form.user_id" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="系统" prop="os" borderBottom>
        <uv-input v-model="form.os" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="浏览器" prop="browser" borderBottom>
        <uv-input v-model="form.browser" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="IP地址" prop="ip" borderBottom>
        <uv-input v-model="form.ip" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="屏幕宽度" prop="width" borderBottom>
        <uv-input v-model.number="form.width" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="屏幕高度" prop="height" borderBottom>
        <uv-input v-model.number="form.height" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="ua记录" prop="ua" borderBottom>
        <uv-input v-model="form.ua" border="surround"></uv-input>
      </uv-form-item>

      <uv-button
        type="primary"
        text="提交"
        customStyle="margin: 40rpx 0"
        @click="submit"
      ></uv-button>
    </uv-form>
  </view>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import {
  monitor_client_detail,
  monitor_client_edit,
  monitor_client_add,
} from "@/api/monitor_client";
import type { type_monitor_client_edit } from "@/api/monitor_client";

import { toast, alert } from "@/utils/utils";
import { useDictData } from "@/hooks/useDictOptions";

let formRef = ref();
let form = ref<type_monitor_client_edit>({
  id: undefined,
  project_key: "",
  client_id: "",
  user_id: "",
  os: "",
  browser: "",
  ip: "",
  width: undefined,
  height: undefined,
  ua: "",
});
const formRules = {
  project_key: [
    {
      required: true,
      message: "请输入项目key",
      trigger: ["blur"],
    },
  ],
  client_id: [
    {
      required: true,
      message: "请输入sdk生成的客户端id",
      trigger: ["blur"],
    },
  ],
  os: [
    {
      required: true,
      message: "请输入系统",
      trigger: ["blur"],
    },
  ],
  browser: [
    {
      required: true,
      message: "请输入浏览器",
      trigger: ["blur"],
    },
  ],
  ua: [
    {
      required: true,
      message: "请输入ua记录",
      trigger: ["blur"],
    },
  ],
};
onLoad((e) => {
  console.log("onLoad", e);
  if (e.id) {
    getDetails(e.id);
  }
});

function getDetails(id) {
  monitor_client_detail(id)
    .then((res) => {
      if (res.code == 200) {
        if (res?.data) {
          form.value = res?.data;
        }
      } else {
        toast(res.message);
      }
    })
    .catch((err) => {
      toast(err.message || "网络错误");
    });
}

function submit() {
  console.log("submit", form.value);
  formRef.value.validate().then(() => {
    if (form.value.id) {
      monitor_client_edit(form.value).then((res) => {
        if (res.code == 200) {
          toast("编辑成功");
          getDetails(form.value?.id);
        } else {
          toast(res.message);
        }
      });
    } else {
      monitor_client_add(form.value).then((res) => {
        if (res.code == 200) {
          toast("添加成功");
          uni.navigateBack();
        } else {
          toast(res.message);
        }
      });
    }
  });
}
</script>

<style lang="scss" scoped>
.page-content {
  padding: 10rpx 20rpx 300rpx;
}
</style>

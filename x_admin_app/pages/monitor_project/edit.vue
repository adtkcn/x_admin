<template>
  <view class="page-content">
    <uv-form
      labelPosition="left"
	  labelWidth="80" 
      :model="form"
      :rules="formRules"
      ref="formRef"
    >
 
      <uv-form-item label="项目uuid" prop="projectKey" borderBottom>
        <uv-input v-model="form.projectKey" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="项目名称" prop="projectName" borderBottom>
        <uv-input v-model="form.projectName" border="surround"></uv-input>
      </uv-form-item>
      <uv-form-item label="项目类型" prop="projectType" borderBottom>
        <!-- <x-picker v-model="form.projectType" valueKey="value" labelKey="name" :columns="dictData.project_type"></x-picker> -->
      </uv-form-item>

      <uv-button
        type="primary"
        text="提交"
        customStyle="margin-top: 20rpx"
        @click="submit"
      ></uv-button>
    </uv-form>
  </view>
</template>

<script setup>
import { reactive, ref, computed } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import {
  monitor_project_detail,
  monitor_project_edit,
} from "@/api/monitor_project";
import { toast, alert } from "@/utils/utils";
import { useDictData } from "@/hooks/useDictOptions";
const { dictData } = useDictData(['project_type'])
let formRef = ref();
let form = ref({
  id: "",
  projectKey: "",
  projectName: "",
  projectType: "",
});
const formRules = {
  projectKey: [
    {
      required: true,
      message: "请输入项目uuid",
      trigger: ["blur"],
    },
  ],
  projectName: [
    {
      required: true,
      message: "请输入项目名称",
      trigger: ["blur"],
    },
  ],
  projectType: [
    {
      required: true,
      message: "请选择项目类型",
      trigger: ["blur"],
    },
  ],
};
onLoad((e) => {
  console.log("onLoad", e);
  getDetails(e.id);
});

function getDetails(id) {
  monitor_project_detail(id)
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

function submit(item) {
  console.log("submit", form);
  formRef.value.validate().then(() => {
    monitor_project_edit(form.value).then((res) => {
      if (res.code == 200) {
        toast("编辑成功");

        getDetails(form.value?.id);
      } else {
        toast(res.message);
      }
    });
  });
}
</script>

<style lang="scss" scoped>
.page-content {
  padding: 10rpx 20rpx 300rpx;
}
</style>

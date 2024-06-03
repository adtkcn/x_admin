<template>
  <view class="page-content">
    <uv-form labelPosition="left" labelWidth="80" :model="form" ref="formRef">
      <uv-form-item label="项目uuid" prop="projectKey" borderBottom>
        <uv-input v-model="form.projectKey"> </uv-input>
      </uv-form-item>
      <uv-form-item label="项目名称" prop="projectName" borderBottom>
        <uv-input v-model="form.projectName"> </uv-input>
      </uv-form-item>
      <uv-form-item
        label="项目类型"
        prop="projectType"
        borderBottom
      >
	  <x-picker v-model="form.projectType" valueKey="value" labelKey="name" :columns="dictData.project_type"></x-picker>
      </uv-form-item>
      <uv-form-item label="创建时间" prop="createTime" borderBottom>
        <xDateRange
          v-model:startTime="form.createTimeStart"
          v-model:endTime="form.createTimeEnd"
        ></xDateRange>
      </uv-form-item>
      <uv-form-item label="更新时间" prop="updateTime" borderBottom>
        <xDateRange
          v-model:startTime="form.updateTimeStart"
          v-model:endTime="form.updateTimeEnd"
        ></xDateRange>
      </uv-form-item>

      <uv-button
        type="primary"
        text="搜索"
        customStyle="margin-top: 20rpx"
        @click="submit"
      ></uv-button>
    </uv-form>
  </view>
</template>

<script setup>
import { onLoad } from "@dcloudio/uni-app";
import { reactive, ref, computed } from "vue";
import { toPath, toast, clearObjEmpty } from "@/utils/utils";
import { useDictData } from "@/hooks/useDictOptions";
import xDateRange from "@/components/x-date-range/x-date-range.vue"
const { dictData } = useDictData(['project_type'])
let formRef = ref();
let form = ref({
  projectKey: "",
  projectName: "",
  projectType: "",
  createTimeStart: "",
  createTimeEnd: "",
  updateTimeStart: "",
  updateTimeEnd: "",
});

function submit() {
  console.log("submit", form.value);

  const search = clearObjEmpty(form.value);

  if (Object.keys(search).length === 0) {
    return toast("请输入查询条件");
  }

  toPath("/pages/monitor_project/index", search);
}
</script>

<style lang="scss" scoped>
.page-content {
  padding: 10rpx 20rpx 300rpx;
}
</style>

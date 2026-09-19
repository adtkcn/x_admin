<template>
  <view class="page-content">
    <uv-form labelPosition="left" labelWidth="80" :model="form" >
      <uv-form-item label="项目uuid" prop="project_key" borderBottom>
        <uv-input v-model="form.project_key"> </uv-input>
      </uv-form-item>
      <uv-form-item label="项目名称" prop="project_name" borderBottom>
        <uv-input v-model="form.project_name"> </uv-input>
      </uv-form-item>
      <uv-form-item label="项目类型" prop="project_type" borderBottom>
        <x-picker
          v-model="form.project_type"
          valueKey="value"
          labelKey="name"
          :columns="dictData.project_type"
        ></x-picker>
      </uv-form-item>
      <uv-form-item label="创建时间" prop="create_time" borderBottom>
        <xDateRange
          v-model:startTime="form.create_time_start"
          v-model:endTime="form.create_time_end"
        ></xDateRange>
      </uv-form-item>
      <uv-form-item label="更新时间" prop="update_time" borderBottom>
        <xDateRange
          v-model:startTime="form.update_time_start"
          v-model:endTime="form.update_time_end"
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

<script setup lang="ts">
import { ref } from "vue";
import { toPath, toast, clearObjEmpty } from "@/utils/utils";
import { useDictData } from "@/hooks/useDictOptions";
import xDateRange from "@/components/x-date-range/x-date-range.vue";
import type {type_monitor_project_query} from "@/api/monitor_project";

const { dictData } = useDictData<{
  project_type: any[];
}>(["project_type"]);
 
let form = ref<type_monitor_project_query>({
  project_key: "",
  project_name: "",
  project_type: "",
  create_time_start: "",
  create_time_end: "",
  update_time_start: "",
  update_time_end: "",
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

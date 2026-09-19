<template>
  <view class="page-content">


    <uv-form labelPosition="left" labelWidth="80" :model="form">
      <uv-form-item label="项目uuid" prop="project_key" borderBottom>
        {{ form.project_key }}
      </uv-form-item>
      <uv-form-item label="项目名称" prop="project_name" borderBottom>
        {{ form.project_name }}
      </uv-form-item>
      <uv-form-item label="项目类型" prop="project_type" borderBottom>
        {{ form.project_type }}
      </uv-form-item>
      <uv-form-item label="创建时间" prop="create_time" borderBottom>
        {{ form.create_time }}
      </uv-form-item>
      <uv-form-item label="更新时间" prop="update_time" borderBottom>
        {{ form.update_time }}
      </uv-form-item>
    </uv-form>
    <uv-button
      v-if="perms('admin:monitor_project:edit')"
      type="primary"
      text="编辑"
      customStyle="margin: 40rpx 0"
      @click="edit"
    ></uv-button>
  </view>
</template>

<script setup lang="ts">
import { reactive, ref, computed } from "vue";
import { onLoad, onShow, onPullDownRefresh } from "@dcloudio/uni-app";

import { useDictData } from "@/hooks/useDictOptions";
import { perms } from "@/utils/perms";
import { monitor_project_detail } from "@/api/monitor_project";
import type { type_monitor_project } from "@/api/monitor_project";

import { toast, alert, toPath } from "@/utils/utils";

let form = ref<type_monitor_project>({
  id: "",
  project_key: "",
  project_name: "",
  project_type: "",
  create_time: "",
  update_time: "",
});
onLoad((e) => {
  console.log("onLoad", e);
  if (e.id) {
    getDetails(e.id);
  }
});
onShow((e) => {
  if (form.value?.id) {
    getDetails(form.value.id);
  }
});
onPullDownRefresh(() => {
  getDetails(form.value.id);
});
function getDetails(id) {
  monitor_project_detail(id)
    .then((res) => {
      uni.stopPullDownRefresh();
      if (res.code == 200) {
        if (res.data) {
          form.value = res?.data;
        }
      } else {
        toast(res.message);
      }
    })
    .catch((err) => {
      uni.stopPullDownRefresh();
      toast(err.message || "网络错误");
    });
}

function edit() {
  toPath("/pages/monitor_project/edit", { id: form.value.id });
}
</script>

<style lang="scss" scoped>
.page-content {
  padding: 10rpx 20rpx 300rpx;
}
</style>
